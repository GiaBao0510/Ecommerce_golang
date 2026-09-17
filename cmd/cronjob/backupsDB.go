package cronjob

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

/*------------ Các tác vụ liên quan đến việc sao lưu DB --------------------*/

const (
	backupTimeout = 30 * time.Minute // pg_dump có thể chạy lâu nếu DB lớn
	uploadTimeout = 10 * time.Minute
)

// hàm RunBackupJob gồm 3 bước theo thứ tự: backup local -> upload cloud -> xóa các bản backup cũ
func RunBackupJob(ctx context.Context, backupDir string, retentionDays int) {
	
	startTime := time.Now()

	// Bước 1: Backup local
	backupCtx, cancelBackup := context.WithTimeout(ctx, backupTimeout)
	filename, err := runBackupLocal(backupCtx, backupDir)
	cancelBackup() // Hủy context sau khi hoàn thành backup local

	if err != nil {
		global.Logger.Error.Error("Lỗi khi sao lưu cơ sở dữ liệu: ", zap.Error(err))
		return
	}
	global.Logger.Access.Info("Sao lưu cơ sở dữ liệu thành công: ", zap.String("file", filename), zap.Duration("duration", time.Since(startTime)))

	// Bước 2: Upload cloud
	uploadCtx, cancelUpload  := context.WithTimeout(ctx, uploadTimeout)
	defer cancelUpload()

	if err := uploadToCloud_R2(uploadCtx, filename); err != nil {
		global.Logger.Warning.Warn("Upload lên Cloudflare R2 thất bại (backup local vẫn an toàn)",
			zap.String("file", filepath.Base(filename)),
			zap.Error(err),
		)
	} else {
		global.Logger.Access.Info("Upload lên Cloudflare R2 thành công",
			zap.String("file", filepath.Base(filename)))
	}

	// Bước 3: Xóa các bản backup cũ
	if err := cleanupOldBackups(backupDir, retentionDays); err != nil {
		global.Logger.Error.Error("Lỗi khi xóa các bản backup cũ: ", zap.Error(err))
	}
}

// hàm runBackupLocal sẽ thực hiện việc sao lưu cơ sở dữ liệu vào thư mục backupDir và trả về đường dẫn của file backup vừa tạo
func runBackupLocal(ctx context.Context, backupDir string) (string, error) {

	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(backupDir, fmt.Sprintf("ecommerce_%s.dump", timestamp))

	cmd := exec.CommandContext(ctx, "pg_dump", // Sử dụng lệnh pg_dump để sao lưu cơ sở dữ liệu PostgreSQL
		"-h", global.Config.PostgreSQL.Host, // Host của cơ sở dữ liệu
		"-p", strconv.Itoa(global.Config.PostgreSQL.Port), // Port của cơ sở dữ liệu
		"-U", global.Config.PostgreSQL.User, // User của cơ sở dữ liệu
		"-d", global.Config.PostgreSQL.DBName, // Tên cơ sở dữ liệu cần sao lưu
		"-F", "c", // Định dạng sao lưu: custom
		"-f", filename, // Đường dẫn file backup
	)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+global.Config.PostgreSQL.Password) // Thêm biến môi trường PGPASSWORD để cung cấp mật khẩu

	output, err := cmd.CombinedOutput() // Thực thi lệnh và lấy output
	if err != nil {
		global.Logger.Error.Error("pg_dump thực thi thất bại",
			zap.String("output", strings.TrimSpace(string(output))),
			zap.Error(err),
		)

		// Để tránh có lỗi file .dump đang dang dở pg_dump thất bại giữa chừng
		//  nếu để lại, cleanupOldBackups sẽ tưởng đó là bản backup hợp lệ
		// (vì vẫn có đuôi .dump), gây nhầm lẫn nghiêm trọng lúc cần restore
		if reomoveErr := os.Remove(filename); reomoveErr != nil && !os.IsNotExist(reomoveErr) {
			global.Logger.Warning.Warn("Không xoá được file backup dang dở",
				zap.String("file", filename), zap.Error(reomoveErr))
		}
		return "", apperrors.NewDetailedInternalServerError("Lỗi khi sao lưu cơ sở dữ liệu", err)
	}

	return filename, nil
}

// Hàm uploadToCloud_R2 sẽ upload file backup lên Cloudflare R2
func uploadToCloud_R2(ctx context.Context, filePath string) error {

	// Load cấu hình AWS từ biến môi trường
	secretKey := global.Config.Authentication.Cloudflare.SecretAccessKey
	bucketName := global.Config.Authentication.Cloudflare.BucketName
	accessKey := global.Config.Authentication.Cloudflare.AccessKeyID
	endpoint := global.Config.Authentication.Cloudflare.Endpoint

	// Kiểm tra các biến môi trường cần thiết
	if endpoint == "" || secretKey == "" || bucketName == "" || accessKey == "" {
		return fmt.Errorf("Missing Cloudflare R2 configuration in environment variables")
	}

	// Tạo cấu hình AWS với thông tin từ Cloudflare R2
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
		config.WithRegion("auto"), // R2 dùng "auto" — không có khái niệm region như AWS thật
	)
	if err != nil {
		global.Logger.Error.Error("không load được cấu hình AWS SDK: ", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	file, err := os.Open(filePath)
	if err != nil {
		global.Logger.Error.Error("không mở được file backup: ", zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}
	defer file.Close()

	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(filePath)),
		Body:   file,
	})
	if err != nil {
		global.Logger.Error.Error("không upload được file backup lên R2: ", zap.Error(err))
		return apperrors.NewDetailedInternalServerError("Không upload được file backup lên R2", err)
	}

	return nil
}

func cleanupOldBackups(backupDir string, retentionDays int) error {
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return apperrors.NewDetailedInternalServerError("Không đọc được thư mục backup", err)
	}

	cutoff := time.Now().AddDate(0, 0, -retentionDays)
	deletedCount := 0

	for _, entry := range entries {

		// Bỏ qua các file không phải là file backup (không có đuôi .dump)
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".dump") {
			continue
		}

		// Lấy thông tin file để kiểm tra thời gian tạo
		info, err := entry.Info()
		if err != nil {
			global.Logger.Error.Error("Không lấy được thông tin file: ", zap.Error(err))
			continue
		}

		// Nếu file backup cũ hơn thời gian cutoff, xóa nó
		if info.ModTime().Before(cutoff) {
			path := filepath.Join(backupDir, entry.Name())
			if err := os.Remove(path); err != nil {
				global.Logger.Error.Error("Không xóa được file backup cũ: ", zap.String("file", path), zap.Error(err))
				continue
			}

			global.Logger.Access.Info("Đã xóa file backup cũ: ", zap.String("file", path))
			deletedCount++
		}
	}

	if deletedCount > 0 {
		global.Logger.Access.Info("Đã dọn dẹp các bản backup cũ",
			zap.Int("deleted_count", deletedCount),
			zap.Int("retention_days", retentionDays),
		)
	}

	return nil
}