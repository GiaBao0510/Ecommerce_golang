package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến sản phẩm của người dùng
func (r *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {

	// Public routes
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
		productRouterPublic.GET("/list")
	}

	// Private routes (cần xác thực)
}
