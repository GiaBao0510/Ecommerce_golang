package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRoleRouter struct{}

func (rl *UserRoleRouter) InitUserRoleRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	rlController, err := wire.InitUserRoleRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo UserRoleRouterHandler: " + err.Error())
	}

	// Private routes
	Router.GET("/user/:uuid", controller.Build(rlController.GetRolesByUserID, logger))
	Router.GET("/role/:id", controller.Build(rlController.GetUserByRoleID, logger))
	Router.POST("", controller.Build(rlController.Create, logger))
	Router.PUT("", controller.Build(rlController.Update, logger))
	Router.DELETE("", controller.Build(rlController.Delete, logger))
}
