package backup

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
	"go.uber.org/zap"
)

// Đây là phần Concrete Strategy, implement BackupStrategy interface để thực hiện backup với Cloudflare
type PostgreSQLBackupStrategy struct {}

// sẽ thực hiện việc sao lưu cơ sở dữ liệu vào thư mục backupDir và trả về đường dẫn của file backup vừa tạo
func (l *PostgreSQLBackupStrategy) Run(
	ctx context.Context, 
	backupDir string,
) (string, error) {

	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(
		backupDir, 
		fmt.Sprintf("ecommerce_%s.dump", timestamp))

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