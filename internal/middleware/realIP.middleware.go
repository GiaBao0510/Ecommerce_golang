package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const RealIPKey = "real_ip"

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
		if ip == "" {
			ip = c.ClientIP() // Nếu không có header nào, lấy IP từ c.ClientIP()
		}

		c.Set(RealIPKey, ip)
		c.Next() // Tiếp tục xử lý request
	}
}
