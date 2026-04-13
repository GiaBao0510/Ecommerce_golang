package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến sản phẩm của người dùng
func (r *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	// Public routes
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.GET("/login")
	}

	// Private routes (cần xác thực)
	adminRouterPrivate := Router.Group("/admin/admin")
	// adminRouterPrivate.Use(Limiter())    // Giới hạn số lượng yêu cầu từ một IP trong một khoảng thời gian nhất định
	// adminRouterPrivate.Use(Authen())     // Middleware xác thực người dùng
	// adminRouterPrivate.Use(Permission()) // Middleware kiểm tra quyền truy cập của người dùng
	{
		adminRouterPrivate.GET("/active_admin")
	}
}