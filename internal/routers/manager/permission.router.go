package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PermissionRouter struct{}

func (r *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	permissionController, err := wire.InitPermissionRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo PermissionRouterHandler: " + err.Error())
	}

	// Private routes - Bọc các router này lại để xác định lỗi tập trung
	Router.GET("", controller.Build(permissionController.GetAll, logger))
	Router.GET("/:id", controller.Build(permissionController.GetByID, logger))
	Router.POST("", controller.Build(permissionController.Create, logger))
	Router.PUT("/:id", controller.Build(permissionController.Update_Put, logger))
	Router.PATCH("/:id", controller.Build(permissionController.Update_Patch, logger))
	Router.DELETE("/:id", controller.Build(permissionController.Delete, logger))
}