package manager

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StatusRouter struct{}

func (r *StatusRouter) InitStatusRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	statusController, err := wire.InitStatusRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo StatusRouterHandler: " + err.Error())
	}

	// Private routes - Bọc các router này lại để xác định lỗi tập trung
	Router.GET("", controller.Build(statusController.GetAllStatuses, logger))
	Router.GET("/:id", controller.Build(statusController.GetStatusByID, logger))
	Router.POST("", controller.Build(statusController.CreateStatus, logger))
	Router.PUT("/:id", controller.Build(statusController.UpdateStatus, logger))
	Router.DELETE("/:id", controller.Build(statusController.DeleteStatus, logger))
}
