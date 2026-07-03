package middleware

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid_token" {
			response.Error_Response(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return 
		}
		c.Next()
	}
}