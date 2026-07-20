package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
	HttpLoggerMiddleware là hàm custom middleware để ghi log cho mỗi HTTP request.

Thì ở đây nó sẽ ghi log:
  - Từ lúc vào - request (method, path, query, body)
  - Từ lúc ra - response (status code, response body, duration)

Tham số: logger *zap.Logger — đây là logger đã được khởi tạo từ pkg/logger.
Middleware này được đăng ký ở router và chạy cho TẤT CẢ các request

	Thông tin được ghi lại:
	  - trace_id: ID duy nhất của request (lấy từ TraceIDMiddleware)
	  - method: GET, POST, PUT, DELETE...
	  - path: /v1/api/statuses/1
	  - client_ip: IP của người dùng gửi request
	  - status_code: 200, 400, 404, 500...
	  - latency_ms: thời gian xử lý (milliseconds)
	  - user_agent: trình duyệt/tool của người dùng (Postman, Chrome,...)
	  - error: nếu có lỗi (chỉ xuất hiện khi status >= 400)
*/
func HttpLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// ---- TRƯỚC KHI XỬ LÝ REQUEST ----

		startTime := time.Now() // Ghi lại thời điểm bắt đầu xử lý request

		ctx.Next() // Thực thi handler tiếp theo (Controller, Service, Repository)

		// ---- SAU KHI XỬ LÝ REQUEST ----

		latency := time.Since(startTime)  // Tính thời gian xử lý
		requestID := GetRequestID(ctx)    // Lấy request_id từ context (được set bởi RequestIDMiddleware)
		realIP, _ := ctx.Get(RealIPKey)   // Lấy real_ip từ context (được set bởi RealIPMiddleware)
		statusCode := ctx.Writer.Status() // Lấy status code trả về cho client

		// Lấy thông tin lỗi nếu có (gin lưu lỗi qua c.Errors)
		// c.Errors là slice chứa các lỗi phát sinh trong quá trình xử lý
		errMsg := ""
		if len(ctx.Errors) > 0 {
			errMsg = ctx.Errors.String()
		}

		// -------------------------------------------------------
		// Tổng hợp các field log chung cho mọi request
		// -------------------------------------------------------
		fields := []zap.Field{
			zap.String("request_id", requestID),               // ID duy nhất của request
			zap.String("real IP", realIP.(string)),            // Real IP của client (nếu có)
			zap.String("method", ctx.Request.Method),          // GET / POST / PUT / DELETE
			zap.String("path", ctx.FullPath()),                // /v1/api/statuses/:id
			zap.String("uri", ctx.Request.RequestURI),         // /v1/api/statuses/1?limit=10
			zap.String("client_ip", ctx.ClientIP()),           // IP của client gửi request
			zap.Int("status_code", statusCode),                // HTTP status code trả về cho client
			zap.Int64("latency_ms", latency.Milliseconds()),   // Thời gian xử lý request (milliseconds)
			zap.String("user_agent", ctx.Request.UserAgent()), // User-Agent của client (Postman, Chrome,...)
		}

		// -------------------------------------------------------
		// Quyết định log ở level nào dựa vào status code
		//
		// 5xx → ERROR (lỗi server, cần điều tra ngay)
		// 4xx → WARN  (lỗi từ phía client, bất thường cần theo dõi)
		// 2xx → INFO  (thành công, thông tin bình thường)
		// -------------------------------------------------------
		if statusCode >= http.StatusInternalServerError {
			// Lỗi server (5xx)
			logger.Error("HTTP Request",
				append(fields, zap.String("error", errMsg))...,
			)
		} else if statusCode >= http.StatusBadRequest {
			// Lỗi client (4xx)
			logger.Warn("HTTP Request",
				append(fields, zap.String("error", errMsg))...,
			)
		} else {
			// Thành công (2xx)
			logger.Info("HTTP Request", fields...)
		}
	}
}
