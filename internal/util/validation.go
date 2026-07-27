package util

import (
	"regexp"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
)

// Hàm kiểm tra các trường bắt buộc trong dữ liệu đầu vào
func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return apperrors.NewBadRequestError(fieldName + " không được để trống")
	}
	return nil
}

// Hàm kiểm tra độ dài của các trường trong dữ liệu đầu vào
func ValidationLength(fieldName, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return apperrors.NewBadRequestError(fieldName + " phải có độ dài từ " + string(min) + " đến " + string(max) + " ký tự")
	}
	return nil
}

// \
func ValidationRegex(fieldname, value string, req *regexp.Regexp) error{
	if
}
