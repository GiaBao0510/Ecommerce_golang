package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// timeoutWriter bọc gin.ResponseWriter, chặn mọi ghi dữ liệu SAU KHI request


// Timeout giới hạn thời gian tối đa xử lý một request,
// tránh một request "treo" chiếm tài nguyên vô thời hạn
// Đây là kiểu middleware đơn giản dựa trên goroutine + select -
// cách này KHÔNG dùng được goroutine xử lý hanler khi timeout xảy ra
func TimeOutMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// 1. Tạo context mới với timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()	// Hủy context khi request kết thúc để tránh rò rỉ goroutine

		// 2. Gắn context mới (có timeout) vào request để truyền tiếp
		c.Request = c.Request.WithContext(ctx)

		// 3. Tạo channel để nhận tín hiệu khi request hoàn tất
		finished := make(chan struct{})

		go func(){
			c.Next() // Tiếp tục xử lý request
			close(finished) // Gửi tín hiệu khi request hoàn tất
		}()

		select {
		case <- finished:
			// Request hoàn tất trước khi timeout
		case <- ctx.Done():
			// Request bị timeout
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"error": "Request timed out",
			})
		}
	}
}