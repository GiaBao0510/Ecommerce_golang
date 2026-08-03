package repositoryimpl

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/lib/pq"
)

// Biến này có vai trò là ánh xạ các trường ràng buộc của PostgreSQL sang các trường tương ứng trong ứng dụng.
// Có thể thêm các ánh xạ cần thiết vào biến này nếu muốn xử lý các lỗi ràng buộc cụ thể.
var pgConstraintFieldMap = map[string]string{
	"USER_email_key":         "email",
	"USER_username_key":      "username",
	"products_name_key":      "name",
	"USER_phone_num_key":     "phone_number",
	"up_review_user_product": "review",
	"uq_oauth_provider_id":   "provider_account",
}

// MapDBError chuyển lỗi thô từ tầng database thành *apperrors.AppError chuẩn hoá.
func MapDBError(err error) error {

	if err == nil {
		return nil
	}

	// 1. Không tìm thấy record - case phổ biến. Nên phải xử lý trước
	if errors.Is(err, sql.ErrNoRows) {
		return apperrors.NewNotFoundError("record not found")
	}

	// 2. Timeout/ context bị hủy (do query chậm hoặc client đóng kết nối)
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return apperrors.NewDatabaseTimeoutError(err)
	}

	//3. Connection pool bị đóng hoặc hết kết nối
	if errors.Is(err, sql.ErrConnDone) || errors.Is(err, driver.ErrBadConn) {
		return apperrors.NewInternalServerError(err)
	}

	// 4. lỗi cụ thể từ postgreSQL
	var pqErr *pq.Error // Tạo một biến con trỏ để kiểm tra lỗi của PostgreSQL
	if !errors.As(err, &pqErr) {
		return apperrors.NewInternalServerError(err)
	}

	switch string(pqErr.Code) {
	case "23505": // Unique violation
		return mapUniqueViolationError(pqErr)
	case "23503": // Foreign key violation
		return apperrors.NewBadRequestError("Dữ liệu tham chiếu không tồn tại hoặc đã bị xóa ")
	case "23502": // Not null violation
		field := pqErr.Column
		if field == "" {
			field = "unknown"
		}
		return apperrors.NewMissingFieldError(field)
	case "23514": // check violation - VD: CHECK (rating between 1 and 5)
		return apperrors.NewInvalidValueError(pqErr.Constraint)
	case "22P02": // invalid_text_representation — VD: UUID/số sai định dạng
		return apperrors.NewInvalidFormatError(pqErr.Column)
	case "08000", "08003", "08006", "08001", "08004", "08007", "08P01": // connection exception
		return apperrors.NewDatabaseConnectionError(err)
	case "57014": // query_canceled - thường do statement timeout hoặc context bị hủy
		return apperrors.NewDatabaseTimeoutError(err)
	case "42P01", "42703": //lỗi này thường do là lỗi code, VD: table/column không tồn tại
		return apperrors.NewInternalServerError(err)
	default:
		return apperrors.NewDatabaseError(err)
	}
}

// mapUniqueViolationError ánh xạ lỗi unique violation từ PostgreSQL sang lỗi ứng dụng tương ứng.
func mapUniqueViolationError(pqErr *pq.Error) error {

	field, ok := pgConstraintFieldMap[pqErr.Constraint]
	if !ok {
		// Constraint chưa được khai báo trong map — fallback an toàn,
		// không lộ Detail. Nên bổ sung constraint này vào pgConstraintFieldMap
		// sau khi xác định được qua log (xem log Error tại tầng gọi).
		return apperrors.NewConflictError("Trường " + field + " đã tồn tại")
	}

	switch pqErr.Constraint {
	case "email", "USER_email_key":
		return apperrors.NewEmailDuplicateError()
	case "phone_number", "USER_phone_num_key":
		return apperrors.NewPhoneDuplicateError()
	default:
		return apperrors.NewConflictError("Trùng lặp dữ liệu" + field)
	}
}

// Hàm này chủ yếu được sử dụng trong các repository
// để ánh xạ lỗi cơ sở dữ liệu sang lỗi ứng dụng, đồng thời thêm thông tin ngữ cảnh (context) vào lỗi để dễ dàng debug và log.
func MapDBErrorWithContext(err error, context string) error {

	mapped := MapDBError(err) // Ánh xạ lỗi cơ sở dữ liệu sang lỗi ứng dụng

	var appErr *apperrors.AppError

	// Nếu lỗi không phải là *apperrors.AppError, thì trả về lỗi đã ánh xạ mà không thêm context
	if !errors.As(mapped, &appErr) {
		return mapped
	}

	switch appErr.Code {
	case http.StatusNotFound:
		return apperrors.NewNotFoundError(context)
	case http.StatusBadRequest:
		return apperrors.NewBadRequestError(context)
	case http.StatusConflict:
		if errors.Is(mapped, apperrors.ErrConflict) {
			return apperrors.NewConflictError(context + " đã tồn tại")
		}
		return mapped

	default:
		return mapped // Trả về lỗi đã ánh xạ mà không thêm context nếu không phải các trường hợp trên
	}
}
