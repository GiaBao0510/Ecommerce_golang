package servicesupport

import (
	"strings"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

// RequireNonEmptyString kiểm tra 1 chuỗi bắt buộc không được rỗng.
func RequireNonEmptyString(value, fieldname, operation string, slog *loghelper.ServiceLogger) error {
	if strings.TrimSpace(value) == "" {
		slog.LogValidationFailed(operation, fieldname + " đang bị rỗng hoặc chỉ chứa khoảng trắng")
		return apperrors.NewBadRequestError(fieldname + " là bắt buộc")
	}
	return nil
}

// RequirePositiveID32 kiểm tra ID kiểu int32 phải là số dương — dùng cho
func RequirePositiveID32(id int32, fieldName, operation string, slog *loghelper.ServiceLogger) error {
	if id <= 0 {
		slog.LogValidationFailed(operation, fieldName + " phải là số dương")
		return apperrors.NewBadRequestError(fieldName + " phải là số dương")
	}
	return nil
}

// RequireNonEmptyUUID kiểm tra ID kiểu string (UUID) không được rỗng — dùng
func RequireNonEmptyUUID(uuid, fieldName, operation string, slog *loghelper.ServiceLogger) error {
	if strings.TrimSpace(uuid) == "" {
		slog.LogValidationFailed(operation, fieldName + " đang bị rỗng hoặc chỉ chứa khoảng trắng")
		return apperrors.NewBadRequestError(fieldName + " là bắt buộc")
	}
	return nil
}

func ValidateOrFail(condition bool, reasonForLog, clientMessage, operation string, slog *loghelper.ServiceLogger) error {
	if !condition {
		slog.LogValidationFailed(operation, reasonForLog)
		return apperrors.NewBadRequestError(clientMessage)
	}
	return nil
}