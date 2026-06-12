package user

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến sản phẩm của người dùng
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// Áp dụng tool WIRE, chúng ta sẽ không cần phải khởi tạo thủ công các service và repository nữa, mà sẽ để WIRE tự động tạo ra các instance và inject chúng vào controller khi cần thiết
	userController, err := wire.InitUserRouterHandler()
	if err != nil {
		// fallback nếu có lỗi xảy ra trong quá trình khởi tạo, có thể log lỗi hoặc xử lý theo cách phù hợp
		panic("Lỗi khi khởi tạo UserRouterHandler: " + err.Error())
	}

	// Public routes
	Router.GET("/register", userController.RegisterController)
	Router.POST("/otp", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "otp endpoint chưa implement",
		})
	})

	// Private routes (cần xác thực)
	// UserRouterPrivate.Use(Limiter())    // Giới hạn số lượng yêu cầu từ một IP trong một khoảng thời gian nhất định
	// UserRouterPrivate.Use(Authen())     // Middleware xác thực người dùng
	// UserRouterPrivate.Use(Permission()) // Middleware kiểm tra quyền truy cập của người dùng
	Router.GET("/get_info", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "get_info endpoint chưa implement",
		})
	})
}
