package middleware

import (
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid_token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "Unauthorized")
			c.Abort()
			return 
		}
		c.Next()
	}
}