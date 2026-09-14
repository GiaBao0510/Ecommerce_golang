package middleware

import (
	"net/http"
	"strings"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthenMiddleware struct {
	redis repository.IRedisRepository
}

func NewAuthenMiddleware(redis repository.IRedisRepository) *AuthenMiddleware {
	return &AuthenMiddleware{
		redis: redis,
	}
}

// AuthenMiddleware là một middleware trong Gin framework để xác thực người dùng dựa trên token trong header Authentication.
// Nó kiểm tra xem header có tồn tại và có hợp lệ hay không. Nếu không hợp lệ, nó trả về lỗi 401 Unauthorized và dừng xử lý request. Nếu hợp lệ, nó cho phép request tiếp tục đến handler tiếp theo.
func (m *AuthenMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Kiểm tra token có trong header Authorization
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing or invalid Authorization header",
			})
			return
		}

		// Lấy token từ header Authorization
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing or invalid Authorization header",
			})
			return
		}

		// Parse token và lấy claims
		_, claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired access token",
			})
			return
		}

		// Lấy JTI từ claims để kiểm tra blacklist/whitelist
		jti, ok := claims["jwt_id"].(string)
		if !ok || jti == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired access token",
			})
			return
		}	

		// Kiểm tra JTI của token này nếu xuất hiện trong Blacklist hoặc không có trong whitelist thì trả về lỗi 401 Unauthorized
		isBlacklisted, err := m.redis.Exists(c, _const.BlackList + jti)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		if isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is blacklisted",
			})
			return
		}

		// Giải mã payload từ token để lấy thông tin người dùng
		payload, err := util.DecryptPayloadFromToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired access token",
			})
			return
		}

		// Lưu thông tin người dùng vào context để các handler tiếp theo có thể sử dụng
		c.Set("user_id", payload.UserID)
		c.Set("user_email", payload.Email)
		c.Set("user_role", payload.UserRole)

		c.Next()
	}
}
