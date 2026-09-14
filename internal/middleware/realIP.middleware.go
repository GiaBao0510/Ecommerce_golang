package middleware

import (
	"strings"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/gin-gonic/gin"
)

// RealIPMiddleware xác định IP thật của client, kể cả khi request đi qua
// reverse proxy / load balancer (Nginx, Cloudflare, AWS ALB...).
func RealIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.GetHeader("X-Real-IP") // Lấy IP từ header X-Real-IP (nếu có)
		if ip == "" {
			if forwarded := c.GetHeader("X-Forwarded-For"); forwarded != "" {
				// X-Forwareded-For có thể chứa nhiều IP, ví dụ: "client_ip, proxy1_ip, proxy2_ip". Lấy IP đầu tiên
				ip = strings.TrimSpace(strings.Split(forwarded, ",")[0])
			}
		}

		// Nếu vẫn không có IP từ header, lấy IP từ c.ClientIP()
		if ip == "" {
			ip = c.ClientIP() // Nếu không có header nào, lấy IP từ c.ClientIP()
		}

		// Nếu vẫn không có IP, lấy IP từ c.Request.RemoteAddr (Trường hợp do bên Client có dùng Proxy để fake IP)
		if ip == "" {
			ip = c.Request.RemoteAddr
		}

		c.Set(_const.RealIPKey, ip)
		c.Next() // Tiếp tục xử lý request
	}
}
