package middleware

import (
	"github.com/gin-gonic/gin"
)

const RequestIDKey = "request_id" // Key để lưu request ID vào context của Gin


// hàm này "X-Request-ID" chủ yếu để lấy tương thích cho các hệ thống/ API Gateway bên ngoài dùng đến tên Header này
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		traceID := GetTraceID(c)	// Lấy trace_id từ context nếu đã có
		// 3. Lưu trace_id vào gin.Context của request hiện tại
		c.Set(RequestIDKey, traceID)

		// 4. Đặt trace_ID vào trong response header
		c.Header("X-Request-ID", traceID)

		c.Next() // Tiếp tục xử lý request
	}
}
