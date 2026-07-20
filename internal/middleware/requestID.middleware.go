package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id" // Key để lưu request ID vào context của Gin

// RequestID gắn UUID duy nhất cho mỗi request — dùng để trace log,
// liên kết log/metric/trace của cùng một request xuyên suốt hệ thống
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Lấy trace_id từ header "X-Request-ID"
		requestID := c.GetHeader("X-Request-ID")

		// 2. Nếu không có trace_id từ client → tự sinh UUID mới
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 3. Lưu trace_id vào gin.Context của request hiện tại
		c.Set(RequestIDKey, requestID)

		// 4. Đặt trace_ID vào trong response header
		c.Header("X-Request-ID", requestID)

		c.Next() // Tiếp tục xử lý request
	}
}

// GetRequestID là hàm helper để lấy request_id từ context một cách an toàn.
//
// Lý do cần hàm này:
// → Thay vì phải viết c.GetString("request_id") ở mọi nơi (dễ typo),
//
//	ta tập trung logic lấy request_id vào một chỗ.
//
// Cách dùng trong các tầng khác:
//
//	requestID := middleware.GetRequestID(c)
func GetRequestID(c *gin.Context) string {
	requestID, exists := c.Get(RequestIDKey)
	if !exists {
		return "" // Nếu không tìm thấy request_id nào thì trả về rỗng
	}

	// type assertion để đảm bảo requestID là string
	id, ok := requestID.(string)
	if !ok {
		return "" // Nếu không phải string thì trả về rỗng
	}

	return id
}
