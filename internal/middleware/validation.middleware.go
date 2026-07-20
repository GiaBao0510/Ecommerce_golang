package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateBody bind và validate request body theo struct tag,
// đặt SÁT TRƯỚC business handler — request không hợp lệ bị chặn
func ValidationMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Validation failed",
				"details": err.Error(),
			})
			return
		}

		c.Set("validated_body", model)
		c.Next()
	}
}
