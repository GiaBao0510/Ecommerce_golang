package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến sản phẩm của người dùng
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// Public routes
	Router.GET("/register")
	Router.POST("/otp")

	// Private routes (cần xác thực)
	// UserRouterPrivate.Use(Limiter())    // Giới hạn số lượng yêu cầu từ một IP trong một khoảng thời gian nhất định
	// UserRouterPrivate.Use(Authen())     // Middleware xác thực người dùng
	// UserRouterPrivate.Use(Permission()) // Middleware kiểm tra quyền truy cập của người dùng
	Router.GET("/get_info")
}
