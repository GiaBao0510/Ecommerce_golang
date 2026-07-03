package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ContextKey là kiểu dùng để đặt key trong context.
// Tại sao không dùng string trực tiếp?
// → Tránh xung đột key nếu các package khác cũng dùng key cùng tên.
type ContextKey string

// TraceIDKey là key để lấy trace_id ra từ context.
// Khai báo constant để dùng ở nhiều nơi mà không sợ typo.
const traceIDKey ContextKey = "trace_id"

// TraceIDMiddleware là middleware chạy ĐẦU TIÊN cho mỗi request.
//
// Nhiệm vụ:
//   1. Kiểm tra header "X-Trace-ID" xem client có gửi trace_id lên không
//      (trường hợp API Gateway hoặc frontend tự sinh trace_id)
//   2. Nếu không có → tự sinh UUID mới
//   3. Lưu trace_id vào gin.Context để các tầng sau (controller, service, repo) dùng chung
//   4. Đặt trace_id vào response header để client có thể dùng khi báo lỗi
func TraceID_Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 1. Lấy trace_id từ header "X-Trace-ID"
		traceID := ctx.GetHeader("X-Trace-ID")

		//2. Nếu không có trace_id từ client → tự sinh UUID mới
		if traceID == ""{
			traceID = uuid.New().String()
		}

		// 3. Lưu trace_id vào gin.Conttext của request hiện tại
		// Từ đây, bất kỳ tầng (Controller, Service, Repository) nào cũng có thể lấy trace_id ra dùng
		ctx.Set(string(traceIDKey), traceID)

		//4. Đặt trace_ID vào trong response header
		// Khi frontend nhận được lỗi, họ có thể dùng trace_id này để báo với bạn
        // Chúng ta vào log filter theo trace_id là thấy ngay toàn bộ hành trình
		ctx.Header("X-Trace-ID", traceID)

		// 5. Tiếp tục xử lý request
		ctx.Next()
	}
}

// GetTraceID  là hàm helper để lấy trace_id từ context một cách an toàn.
//
// Lý do cần hàm này:
// → Thay vì phải viết c.GetString("trace_id") ở mọi nơi (dễ typo),
//   ta tập trung logic lấy trace_id vào một chỗ.
//
// Cách dùng trong các tầng khác:
//   traceID := middleware.GetTraceID(c)
func GetTraceID(c *gin.Context) string {
	traceID, exists := c.Get(string(traceIDKey))
	if !exists {
		return "" // Nếu không tìm thấy trace_id nào thì trả về rỗng
	}

	// type assertion để đảm bảo traceID là string
	id, ok := traceID.(string)
	if !ok {
		return "" // Nếu không phải string thì trả về rỗng
	}

	return id
}