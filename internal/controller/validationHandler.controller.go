package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/go-playground/validator/v10"
)

// FieldErrorDetail là chi tiết lỗi cho từng trường, bao gồm tên trường và thông báo lỗi
type FieldErrorDetail struct {
	Field   string `json:"field"`   // Tên trường bị lỗi
	Message string `json:"message"` // Thông báo lỗi
}

// HandleValidationError xử lý lỗi từ tầng BINDING/VALIDATION (trước khi vào logic)
func HandleValidationError(err error) *apperrors.AppError {

	// 1. Body rỗng hoặc không hợp lệ
	if errors.Is(err, io.EOF) {
		return apperrors.NewBadRequestError("Body rỗng hoặc không hợp lệ")
	}

	// 2. Lỗi sai kiểu dữ liệu khi unmarshal JSON (ví dụ: gửi chuỗi thay vì số)
	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return apperrors.NewBadRequestError(fmt.Sprintf("Sai kiểu dữ liệu cho trường '%s', giá trị nhận được: '%v'", ute.Field, ute.Value))
	}

	//3. Lỗi sai cú pháp
	var sy *json.SyntaxError
	if errors.As(err, &sy) {
		return apperrors.NewBadRequestError(fmt.Sprintf("Sai cú pháp JSON tại byte offset %d: %v", sy.Offset, sy.Error()))
	}

	// 4. Kiểm tra xem lỗi có phải từ validator không
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {

		// 2. Duyệt qua từng trường bị lỗi và tạo thông báo
		details := make([]string, 0, len(ve))
		for _, fe := range ve {
			details = append(details, translateFieldError(fe))
		}

		// 4. Trả về lỗi ứng dụng với thông báo chi tiết
		return apperrors.NewBadRequestError(strings.Join(details, ";"))
	}

	// Nếu không phải lỗi validator, trả về lỗi gốc
	return apperrors.NewBadRequestError("Dữ liệu không hợp lệ: " + err.Error())
}

// Hàm này dịch thông báo lỗi từ validator sang thông báo dễ hiểu hơn
// validator.FieldError cung cap:
//
//	.Field()     → Ten truong struct (VD: "Name", "Email")
//	.Tag()       → Ten rule validation (VD: "required", "min", "max")
//	.Param()     → Tham so cua rule (VD: "2" trong min=2, "100" trong max=100)
//	.Value()     → Gia tri thuc te ma client da gui
//	.Namespace() → Duong dan day du (VD: "CreateStatusRequest.Name")
func translateFieldError(fe validator.FieldError) string {
	field := fe.Field()
	isNumeric := fe.Kind() >= reflect.Int && fe.Kind() <= reflect.Float64

	// Dựa vào rule validation, tạo thông báo lỗi phù hợp
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s: trường này bắt buộc", field)
	case "min":
		if isNumeric {
			return fmt.Sprintf("%s: giá trị phải lớn hơn hoặc bằng %s", field, fe.Param())
		}
		return fmt.Sprintf("%s: giá trị quá ngắn, tối thiểu là %s", field, fe.Param())
	case "max":
		if isNumeric {
			return fmt.Sprintf("%s: giá trị phải nhỏ hơn hoặc bằng %s", field, fe.Param())
		}
		return fmt.Sprintf("%s: giá trị quá dài, tối đa là %s", field, fe.Param())
	case "gt":
		return fmt.Sprintf("%s: giá trị phải lớn hơn %s", field, fe.Param())
	case "lt":
		return fmt.Sprintf("%s: giá trị phải nhỏ hơn %s", field, fe.Param())
	case "gte":
		return fmt.Sprintf("%s: giá trị phải lớn hơn hoặc bằng %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("%s: giá trị phải nhỏ hơn hoặc bằng %s", field, fe.Param())
	case "email":
		return fmt.Sprintf("%s: định dạng Email không hợp lệ", field)
	case "uuid":
		return fmt.Sprintf("%s: định dạng UUID không hợp lệ", field)
	case "oneof":
		return fmt.Sprintf("%s: giá trị phải là một trong các giá trị hợp lệ: %s", field, fe.Param())
	case "e164":
		return fmt.Sprintf("%s: định dạng số điện thoại không hợp lệ", field)
	case "len":
		if isNumeric {
			return fmt.Sprintf("%s: độ dài phải là %s", field, fe.Param())
		}
		return fmt.Sprintf("%s: độ dài phải là %s ký tự", field, fe.Param())
	case "numeric":
		return fmt.Sprintf("%s: giá trị phải là số", field)
	case "url":
		return fmt.Sprintf("%s: định dạng URL không hợp lệ", field)
	case "datetime":
		return fmt.Sprintf("%s: định dạng ngày giờ không hợp lệ, phải theo định dạng %s", field, fe.Param())
	case "alpha":
		return fmt.Sprintf("%s: giá trị chỉ được chứa các ký tự chữ cái", field)
	default:
		return fmt.Sprintf("%s: giá trị không hợp lệ", field)
	}
}
