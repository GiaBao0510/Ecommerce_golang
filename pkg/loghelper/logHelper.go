package loghelper

import "go.uber.org/zap"

// Tạo ra struct hỗ trợ cho việc ghi log trên các repository
// Thay vì mỗi Repository struct định nghĩa logDBError() của riêng mình,
type DBLogger struct {
	logger   *zap.Logger // Logger chính để ghi
	repoName string      // Tên repository để phân biệt log từ các repository khác nhau
}

// Hàm khởi tạo
func NewDBLogger(logger *zap.Logger, repoName string) *DBLogger {
	return &DBLogger{
		logger:   logger,
		repoName: repoName,
	}
}

// Ghi log lỗi với thông tin chi tiết
// ERROR: Repository: GetByID failed, layer: repository, operation: GetStatusByID, id: 123, error: sql: no rows in result set
func (d *DBLogger) LogError (operation string, err error, extrafields ...zap.Field) {
	fields := []zap.Field {
		zap.String("repo", d.repoName),
		zap.String("layer", "repository"),
		zap.String("operation", operation),
		zap.Error(err),
	}
	d.logger.Error("Operation failed", append(fields, extrafields...)...)
}

// Ghi log cảnh báo với thông tin chi tiết
func (d *DBLogger) LogWarning (operation string, msg string, extraFields ...zap.Field) {
	fields := []zap.Field {
		zap.String("repo", d.repoName),
		zap.String("layer", "repository"),
		zap.String("operation", operation),
	}
	d.logger.Warn(msg, append(fields, extraFields...)...)
}

