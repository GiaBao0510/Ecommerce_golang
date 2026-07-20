package middleware

import "github.com/gin-gonic/gin"

//Security headers giúp giảm một số rủi ro phổ biến trên trình duyệt.
func SecurityHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload") // Bảo vệ khỏi các cuộc tấn công man-in-the
		c.Header("X-Content-Type-Options", "nosniff")                                         // Ngăn chặn trình duyệt đoán loại nội dung
		c.Header("X-Frame-Options", "DENY")                                                   // Ngăn chặn clickjacking
		c.Header("X-XcSS-Protection", "1; mode=block")                                         // Bật bộ lọc XSS của trình duyệt
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")                        // Kiểm soát thông tin referrer được gửi đi
		c.Next()
	}
}
