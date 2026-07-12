package manager

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RolesRouter struct{}

func (r *RolesRouter) InitRolesRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	rolesController, err := wire.InitRolesRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo RolesRouterHandler: " + err.Error())
	}

	// Private routes - Bọc các router này lại để xác định lỗi tập trung
	Router.GET("", controller.Build(rolesController.GetAll, logger))
	Router.GET("/:id", controller.Build(rolesController.GetByID, logger))
	Router.POST("", controller.Build(rolesController.Create, logger))
	Router.PUT("/:id", controller.Build(rolesController.Update_Put, logger))
	Router.PATCH("/:id", controller.Build(rolesController.Update_Patch, logger))
	Router.DELETE("/:id", controller.Build(rolesController.Delete, logger))
}
