package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracing tạo một span cho mỗi request — hỗ trợ theo dõi request
// xuyên suốt nhiều microservice (ví dụ: order-service → payment-service)
func TracingMiddleware() gin.HandlerFunc {
	
	// 1. Tạo tracer để tạo span
	tracer := otel.Tracer("ecommerce-api")

	return func(c *gin.Context) {

		// 2. Tạo span mới cho request hiện tại
		ctx, span := tracer.Start(c.Request.Context(), c.Request.URL.Path,
		trace.WithAttributes(),)

		defer span.End() // Kết thúc span khi request hoàn tất

		// Gắn context mới (có chứa span) vào request để truyền tiếp
        // xuống các lớp service/repository phía sau
		c.Request = c.Request.WithContext(ctx)
		c.Next() // Tiếp tục xử lý request
	}
}