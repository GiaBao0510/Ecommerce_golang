package manager

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
)

type StatusRouter struct{}

func (r *StatusRouter) InitStatusRouter(Router *gin.RouterGroup, db *database.Queries) {
	statusController, err := wire.InitStatusRouterHandler(db)
	if err != nil {
		panic("Lỗi khi khởi tạo StatusRouterHandler: " + err.Error())
	}

	// Private routes
	Router.GET("", statusController.GetAllStatuses)
	Router.GET("/:id", statusController.GetStatusByID)
	Router.POST("", statusController.CreateStatus)
	Router.PUT("/:id", statusController.UpdateStatus)
	Router.DELETE("/:id", statusController.DeleteStatus)
}
