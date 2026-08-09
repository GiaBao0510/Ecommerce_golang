package http

import (
	"errors"

	"github.com/GiaBao0510/Ecommerce_golang/internal/middleware"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"go.uber.org/zap"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
1️⃣  Nhận một AppHandler (hàm controller trả về error)
2️⃣  Thực thi handler – nếu không có lỗi → kết thúc bình thường
3️⃣  Nếu có lỗi → switch-case kiểm tra kiểu lỗi:
     • errors.As(*AppError) → dùng Code và Status từ struct lỗi
     • errors.Is(ErrNotFound) → trả 404
     • errors.Is(ErrUnauthorized) → trả 401
     • default → trả 500 Internal Server Error
4️⃣  Controller KHÔNG BAO GIỜ tự gọi ctx.JSON khi có lỗi
*/

// AppHandler là một hàm controller trả về lỗi thay vì trả về void
type AppHandler func(ctx *gin.Context) error

// Build là wrapper bao quanh handler, đóng vai trò:
//  1. Thực thi handler
//  2. Nếu có lỗi → log lại + trả response lỗi phù hợp
//  3. Đây là ĐIỂM DUY NHẤT xử lý lỗi — controller không cần tự gọi ctx.JSON khi lỗi
//
// Tham số logger *zap.Logger được inject từ ngoài vào (Dependency Injection).
// Điều này giúp dễ test (có thể mock logger) và tránh dùng global variable.
func Build(handler AppHandler, logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 1. Thực thi logic controller
		err := handler(ctx)
		if err == nil {
			return // Không có lỗi, kết thúc bình thường
		}

		// 2. Lấy trace_id từ context để log
		traceID := middleware.GetTraceID(ctx)

		// 3. Phân loại log lỗi để đưa vào log
		// errors.As() kiểm tra xem err (hoặc lỗi bên trong nó) có phải *AppError không
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {

			if appErr.Code >= http.StatusInternalServerError {
				// 5xx: Lỗi server nghiêm trọng - log ERROR
				logger.Error("Application Error",
					zap.String("trace_id", traceID),
					zap.String("method", ctx.Request.Method),
					zap.String("path", ctx.Request.URL.Path),
					zap.Int("status_code", appErr.Code),
					zap.String("error_status", appErr.Status),
					zap.String("error_message", appErr.Message),
					zap.NamedError("root_cause", appErr.Unwrap()),
					zap.Error(err), // Ghi cả stack trace nếu có
				)
			} else if appErr.Code >= http.StatusBadRequest {
				// 4xx: Lỗi client - log WARN
				logger.Warn("Client Error",
					zap.String("trace_id", traceID),
					zap.String("method", ctx.Request.Method),
					zap.String("path", ctx.Request.URL.Path),
					zap.Int("status_code", appErr.Code),
					zap.String("error_message", appErr.Message),
				)
			}

			// Trả vể response với format chuẩn
			ctx.JSON(appErr.Code, response.ErrorResponse{
				Code:    appErr.Code,
				Status:  appErr.Status,
				Message: appErr.Message,
			})
			return
		}

		// -------------------------------------------------------
		// Bước 4: Fallback — các lỗi không phải AppError
		// (Ít gặp nhưng cần xử lý để không crash app)
		// -------------------------------------------------------
		// switch {
		// case errors.Is(err, apperrors.ErrUnauthorized):
		// 	logger.Warn("Unauthorized access", zap.String("trace_id", traceID), zap.String("path", ctx.Request.URL.Path),)
		// 	ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
		// 		Code:    http.StatusUnauthorized,
		// 		Status:  "Unauthorized",
		// 		Message: "Không có quyền truy cập - detail: " + err.Error(),
		// 	})
		// case errors.Is(err, apperrors.ErrForbidden):
		// 	logger.Warn("Forbidden access", zap.String("trace_id", traceID), zap.String("path", ctx.Request.URL.Path))
		// 	ctx.JSON(http.StatusForbidden, response.ErrorResponse{
		// 		Code:    http.StatusForbidden,
		// 		Status:  "Forbidden",
		// 		Message: "Hành động bị cấm - detail: " + err.Error(),
		// 	})
		// case errors.Is(err, apperrors.ErrNotFound):
		// 	logger.Warn("Not Found", zap.String("trace_id", traceID), zap.String("path", ctx.Request.URL.Path))
		// 	ctx.JSON(http.StatusNotFound, response.ErrorResponse{
		// 		Code:    http.StatusNotFound,
		// 		Status:  "Not Found",
		// 		Message: "Không tìm thấy tài nguyên - detail: " + err.Error(),
		// 	})
		// case errors.Is(err, apperrors.ErrInvalidFormat):
		// 	logger.Warn("Invalid Format", zap.String("trace_id", traceID), zap.String("path", ctx.Request.URL.Path))
		// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
		// 		Code:    http.StatusBadRequest,
		// 		Status:  "Bad Request",
		// 		Message: "Định dạng không hợp lệ - detail: " + err.Error(),
		// 	})
		// case errors.Is(err, apperrors.ErrPhoneDuplicate):
		// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
		// 		Code:    http.StatusBadRequest,
		// 		Status:  "Bad Request",
		// 		Message: "Số điện thoại đã tồn tại - detail: " + err.Error(),
		// 	})
		// default:
		// 	logger.Warn("Forbidden access", 
		// 		zap.String("trace_id", traceID), 
		// 		zap.String("method", ctx.Request.Method),
		// 		zap.String("path", ctx.Request.URL.Path),
		// 		zap.String("error_message", err.Error()),
		// 		zap.Error(err),
		// 	)
		// 	ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
		// 		Code:    http.StatusInternalServerError,
		// 		Status:  "Internal Server Error",
		// 		Message: "Lỗi máy chủ - detail: ",
		// 	})
		// }

		logger.Error("Unhandled error type - missing AppError wrap in code",
			zap.String("trace_id", traceID),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("error_message", err.Error()),
			zap.Error(err),
		)
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Lỗi máy chủ - detail: " + err.Error(),
		})
	}
}
