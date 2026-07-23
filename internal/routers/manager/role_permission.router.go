package manager

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RolePermissionRouter struct{}

func (rp *RolePermissionRouter) InitRolePermissionRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	rpController, err := wire.InitRolePermissionRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo RolePermissionRouterHandler: " + err.Error())
	}

	// Private routes - Bọc các router này lại để xác định lỗi tập trung
	Router.GET("/role/:id", controller.Build(rpController.GetPermissionsByRoleID, logger))
	Router.GET("/permission/:id", controller.Build(rpController.GetRolesByPermissionID, logger))
	Router.POST("", controller.Build(rpController.Create, logger))
	Router.PUT("/:id", controller.Build(rpController.Update_Put, logger))
	Router.DELETE("/:id", controller.Build(rpController.Delete, logger))
}
