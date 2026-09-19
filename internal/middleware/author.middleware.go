package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {

		// lấy role từ context sau khi mà AuthenMiddleware đã xác thực và lưu vào context
		roleValue, exists := c.Get("user_role")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, 
				gin.H{"error": "Role not found in context"})
			return
		}

		// kiểm tra role có hợp lệ không
		userRole, ok := roleValue.(int)
		if !ok {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, 
				gin.H{"error": "Invalid role type in context"})
			return
		}

		// Duyệt qua các role được phép và kiểm tra xem role của người dùng có nằm trong danh sách đó không
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {	// Nếu tìm thấy role hợp lệ, thì cho phép request tiếp tục
				c.Next()
				return
			}
		}

		// Nếu không tìm thấy role hợp lệ, trả về lỗi 403 Forbidden
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "You don't have permission to access this resource"})
	}
}