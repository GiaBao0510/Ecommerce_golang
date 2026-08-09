package user

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
// Chỉ với vai trò là client nên sẽ thao tác một phần
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	userController, err := wire.InitUserRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo UserRouterHandler: " + err.Error())
	}
	// Public routes

	// Private routes (cần xác thực)
	Router.PUT("/:uuid", controller.Build(userController.Update_Put, logger))
	Router.PATCH("/:uuid", controller.Build(userController.Update_Patch, logger))
	Router.DELETE("/:uuid", controller.Build(userController.Delete, logger))
}
