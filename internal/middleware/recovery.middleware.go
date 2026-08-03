package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryMiddleware bắt mọi panic xảy ra từ các middleware/handler phía sau,
// tránh crash toàn bộ server chỉ vì 1 request lỗi. Log panic đầy đủ stack
// trace vào global.Logger.Error (structured, đi vào file log + rotation),
// thay vì chỉ in ra stderr như gin.Recovery() mặc định.
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if rec := recover(); rec != nil {
				traceID := GetTraceID(c) // Lấy trace_id từ context để log


				global.Logger.Error.Error(
					"Panic recovered in middleware",
					zap.String("trace_id", traceID),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Any("panic", rec),
					zap.String("stack_trace", string(debug.Stack())),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   http.StatusInternalServerError,
					"status": "Internal Server Error",
					"message": "Lỗi không mong muốn xảy ra. Vui lòng thử lại sau.",
				})
			}
		}()

		c.Next() // Tiếp tục xử lý request
	}
}