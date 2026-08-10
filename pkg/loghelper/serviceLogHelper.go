package loghelper

import "go.uber.org/zap"

// Tạo ra struct hỗ trợ cho việc ghi log trên các service
type ServiceLogger struct {
	logger      *zap.Logger
	serviceName string
}

// Khởi tạo ServiceLogger cho một service cụ thể, giúp phân biệt log từ các service khác nhau
func NewServiceLogger(logger *zap.Logger, serviceName string) *ServiceLogger {
	return &ServiceLogger{
		logger:      logger,
		serviceName: serviceName,
	}
}

// LogValidationFailed ghi log khi input KHÔNG vượt qua validate nghiệp vụ
func (s *ServiceLogger) LogValidationFailed(operation, reason string, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("service", s.serviceName),
		zap.String("layer", "service"),
		zap.String("operation", operation),
		zap.String("reason", reason),
	}

	s.logger.Warn("Validation failed", append(fields, extraFields...)...)
}

// LogInfo ghi log thông tin nghiệp vụ bình thường (VD: tạo/cập nhật thành công).
func (s *ServiceLogger) LogInfo(operation, message string, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("service", s.serviceName),
		zap.String("layer", "service"),
		zap.String("operation", operation),
	}
	s.logger.Info(message, append(fields, extraFields...)...)
}

// LogError ghi log khi có lỗi xảy ra trong service
func (s *ServiceLogger) LogError(operation string, err error, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("service", s.serviceName),
		zap.String("layer", "service"),
		zap.String("operation", operation),
		zap.Error(err),
	}
	s.logger.Error("Operation failed", append(fields, extraFields...)...)
}

func (s *ServiceLogger) LogWarning(operation, message string, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("service", s.serviceName),
		zap.String("layer", "service"),
		zap.String("operation", operation),
	}
	s.logger.Warn(message, append(fields, extraFields...)...)
}
