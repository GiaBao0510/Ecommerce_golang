package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRouter struct {
}

// Đây là nơi bạn sẽ định nghĩa các route liên quan đến người dùng
// Và với vai trò admin sẽ được phép truy cập tất cả các route này, còn với vai trò user sẽ chỉ được phép truy cập một số route nhất định
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	userController, err := wire.InitUserRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo UserRouterHandler: " + err.Error())
	}
	// Public routes
	// Router.GET("/register")
	// Router.POST("/otp")

	// Private routes (cần xác thực)
	Router.GET("", controller.Build(userController.GetAll, logger))
	Router.GET("/:uuid", controller.Build(userController.GetByID, logger))
	Router.GET("/email/:email", controller.Build(userController.GetUserByEmail, logger))
	Router.GET("/phone/:phone", controller.Build(userController.GetUserByPhone, logger))
	Router.POST("", controller.Build(userController.Create, logger))
	Router.PUT("/:uuid", controller.Build(userController.Update_Put, logger))
	Router.PATCH("/:uuid", controller.Build(userController.Update_Patch, logger))
	Router.DELETE("/:uuid", controller.Build(userController.Delete, logger))
}
