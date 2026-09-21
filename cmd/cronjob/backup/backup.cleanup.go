package backup

import (
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/GiaBao0510/Ecommerce_golang/global"
    "github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
    "go.uber.org/zap"
)

// cleanupOldBackups sẽ xóa các bản backup cũ hơn retentionDays ngày trong thư mục backupDir
func cleanupOldBackups(
	backupDir string, 
	retentionDays int,
) error {
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
				global.Logger.Error.Error("Không xoá được bản backup cũ",
					zap.String("file", path), zap.Error(err))
				continue
			}
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