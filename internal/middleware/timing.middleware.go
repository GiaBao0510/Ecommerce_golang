package middleware

import (
	"github.com/GiaBao0510/Ecommerce_golang/pkg/timing"
	"github.com/gin-gonic/gin"
)

func Timing_Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqCtx, _ := timing.NewCollector(c.Request.Context())
		c.Request = c.Request.WithContext(reqCtx)
		c.Next()
	}
}