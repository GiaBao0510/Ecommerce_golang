package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery bắt mọi panic xảy ra từ các middleware/handler phía sau
// Tránh crash toàn bộ server và chỉ vì request lỗi
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				global.Logger.Error.Error(fmt.Sprintf("[Panic recovered]: %v\n%s\n", err, debug.Stack), zap.Any("error", err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next() // Tiếp tục xử lý request
	}
}