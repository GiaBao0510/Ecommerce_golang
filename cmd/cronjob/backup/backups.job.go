package backup

import (
	"context"
    "path/filepath"
    "time"

    "github.com/GiaBao0510/Ecommerce_golang/global"
    "go.uber.org/zap"
)

// Ở đây sẽ áp dụng triển khai sao lưu từ Cloudflare và Local dựa trên Strategy Pattern
const (
	backupTimeout = 30 * time.Minute // pg_dump có thể chạy lâu nếu DB lớn
	uploadTimeout = 10 * time.Minute
)

type BackupJob struct {
	localContext *BackupContext
	cloudContext *BackupContext
}

func NewBackupJob(
	localStrategy BackupStrategy,
	cloudStrategy BackupStrategy,
) *BackupJob {
	return &BackupJob{
		localContext: NewBackupContext(localStrategy),
		cloudContext: NewBackupContext(cloudStrategy),
	}
}

// hàm RunBackupJob gồm 3 bước theo thứ tự: backup local -> upload cloud -> xóa các bản backup cũ
func(job *BackupJob) Run(
	ctx context.Context, 
	backupDir string, 
	retentionDays int,
) error {
	
	startTime := time.Now()

	backupCtx, cancelBackup := context.WithTimeout(ctx, backupTimeout)
	defer cancelBackup() // Hủy context sau khi hoàn thành backup local

	// Bước 1: Backup local
	filename, err := job.localContext.Run(backupCtx, backupDir)
	if err != nil {
		global.Logger.Error.Error("Lỗi khi sao lưu cơ sở dữ liệu vào local: ", zap.Error(err))
		return err
	}

	global.Logger.Access.Info(
		"Sao lưu cơ sở dữ liệu thành công: ", 
		zap.String("file", filename), 
		zap.Duration("duration", time.Since(startTime)),
	)

	// Bước 2: Upload cloud
	uploadCtx, cancelUpload  := context.WithTimeout(ctx, uploadTimeout)
	defer cancelUpload()

	if _, err := job.cloudContext.Run(uploadCtx, filename); err != nil {
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

	return nil
}

func RunBackupJob(
    ctx context.Context,
    backupDir string,
    retentionDays int,
) {
    job := NewBackupJob(
        &PostgreSQLBackupStrategy{},
        &CloudflareBackupStrategy{},
    )

    if err := job.Run(ctx, backupDir, retentionDays); err != nil {
        global.Logger.Error.Error(
            "Backup job thất bại",
            zap.Error(err),
        )
    }
}

