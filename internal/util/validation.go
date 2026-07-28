package util

import (
	"regexp"
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/google/uuid"
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

// hàm này dùng để kiểm tra xem một chuỗi có khớp với một biểu thức chính quy (regex) hay không. Nếu không khớp, nó sẽ trả về một lỗi BadRequestError với thông báo rằng trường đó không hợp lệ.
func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error{
	if !reg.MatchString(value) {
		return apperrors.NewBadRequestError(fieldname + " không hợp lệ")
	}
	return nil
}

// Hàm kiểm tra xem một chuỗi có phải là một số nguyên dương hay không. Nếu không, nó sẽ trả về một lỗi BadRequestError với thông báo rằng trường đó phải là một số nguyên dương.
func ValidationPositiveInt(fieldName, value string) (int32, error) {
	v ,err := strconv.Atoi(value)
	if err != nil || v <= 0 {
		return 0, apperrors.NewBadRequestError(fieldName + " phải là một số nguyên dương")
	}

	return int32(v), nil
}

// Validation UUID
func ValidationUUID(fieldName, value string) error {
	uid, err := uuid.Parse(value)
	if err != nil || uid == uuid.Nil { 
		return apperrors.NewBadRequestError(fieldName + " không hợp lệ")
	}

	return nil
}