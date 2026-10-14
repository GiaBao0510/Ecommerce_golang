package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryRouter struct{}

func (r *CategoryRouter) InitCategoryRouter(router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	categoryController, err := wire.InitCategoryRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo CategoryRouterHandler: " + err.Error())
	}

	router.GET("", controller.Build(categoryController.GetAll, logger))
	router.GET("/name/:name", controller.Build(categoryController.GetByName, logger))
	router.GET("/:id", controller.Build(categoryController.GetByID, logger))
	router.POST("", controller.Build(categoryController.Create, logger))
	router.PUT("/:id", controller.Build(categoryController.Update_Put, logger))
	router.DELETE("/:id", controller.Build(categoryController.Delete, logger))
}
