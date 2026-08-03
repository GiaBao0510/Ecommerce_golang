package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Tracing tạo một span cho mỗi request — hỗ trợ theo dõi request
// xuyên suốt nhiều microservice (ví dụ: order-service → payment-service)
func TracingMiddleware() gin.HandlerFunc {
	
	// 1. Tạo tracer để tạo span
	tracer := otel.Tracer("ecommerce-api")

	return func(c *gin.Context) {

		// 2. Tạo span mới cho request hiện tại
		ctx, span := tracer.Start(
			c.Request.Context(), 
			c.Request.URL.Path,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.target", c.Request.URL.Path),
			),
		)
		defer span.End() // Kết thúc span khi request hoàn tất

		// Gắn context mới (có chứa span) vào request để truyền tiếp
        // xuống các lớp service/repository phía sau
		c.Request = c.Request.WithContext(ctx)
		c.Next() // Tiếp tục xử lý request

		// 3. Sau khi request hoàn tất, ghi thêm thông tin vào span
		statusCode := c.Writer.Status()
		span.SetAttributes(attribute.Int("http.status_code", statusCode))

		if statusCode >= 500 {
			span.SetStatus(codes.Error, "Internal Server Error")
		}
		if len(c.Errors) > 0 {
			span.RecordError(errors.New(c.Errors.String()))
		}
 	}
}