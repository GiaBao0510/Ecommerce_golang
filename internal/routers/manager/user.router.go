package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến sản phẩm của người dùng
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// Public routes
	UserRouterPublic := Router.Group("/User")
	{
		UserRouterPublic.GET("/register")
		UserRouterPublic.POST("/otp")
	}

	// Private routes (cần xác thực)
	UserRouterPrivate := Router.Group("/user")
	// UserRouterPrivate.Use(Limiter())    // Giới hạn số lượng yêu cầu từ một IP trong một khoảng thời gian nhất định
	// UserRouterPrivate.Use(Authen())     // Middleware xác thực người dùng
	// UserRouterPrivate.Use(Permission()) // Middleware kiểm tra quyền truy cập của người dùng
	{
		UserRouterPrivate.GET("/get_info")
	}
}