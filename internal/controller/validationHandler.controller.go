package controller

import (
	"errors"
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/go-playground/validator/v10"
)

// Hàm này xử lý lỗi vadidation tập trung
// Nó nhận vào một lỗi và trả về một lỗi ứng dụng (AppError) tương ứng
// Hàm này xử lý lỗi từ BINDING/VALIDATION (trước khi đi vào logic)
func HandleValidationError(err error) *apperrors.AppError {

	// 1. Kiểm tra xem lỗi có phải từ validator không
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {

		// 2. Duyệt qua từng trường bị lỗi và tạo thông báo
		details := make([]string,0, len(ve))
		for _, fe := range ve {

			detail := translateFieldError(fe)
			details = append(details, detail)
		}

		// 3. Gom tất cả lỗi thành một chuỗi
		message := ""
		for i, d := range details {
			if i > 0{
				message += "; "
			}
			message += d
		}

		// 4. Trả về lỗi ứng dụng với thông báo chi tiết
		return apperrors.NewBadRequestError(message)
	}

	// Nếu không phải lỗi validator, trả về lỗi gốc
	return apperrors.NewBadRequestError("Dữ liệu không hợp lệ: " + err.Error())
}

// Hàm này dịch thông báo lỗi từ validator sang thông báo dễ hiểu hơn
// validator.FieldError cung cap:
//   .Field()     → Ten truong struct (VD: "Name", "Email")
//   .Tag()       → Ten rule validation (VD: "required", "min", "max")
//   .Param()     → Tham so cua rule (VD: "2" trong min=2, "100" trong max=100)
//   .Value()     → Gia tri thuc te ma client da gui
//   .Namespace() → Duong dan day du (VD: "CreateStatusRequest.Name")
func translateFieldError(fe validator.FieldError) string {
	field := fe.Field()

	// Dựa vào rule validation, tạo thông báo lỗi phù hợp
	switch fe.Tag(){
	case "required":
		return fmt.Sprintf("%s: trường này bắt buộc", field)
	case "min":
		return fmt.Sprintf("%s: giá trị quá ngắn, tối thiểu là %s", field, fe.Param())
	case "max":
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
		return fmt.Sprintf("%s: định dạng Email không hợp lệ",field)
	case "uuid":
		return fmt.Sprintf("%s: định dạng UUID không hợp lệ",field)
	case "oneof":
		return fmt.Sprintf("%s: giá trị phải là một trong các giá trị hợp lệ: %s", field, fe.Param())
	case "e164":
		return fmt.Sprintf("%s: định dạng số điện thoại không hợp lệ", field)
	case "len":
		return fmt.Sprintf("%s: độ dài phải là %s ký tự", field, fe.Param())
	default:
		return fmt.Sprintf("%s: giá trị không hợp lệ", field)
	} 
}