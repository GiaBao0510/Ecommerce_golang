package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SubCategoryRouter struct{}

func (r *SubCategoryRouter) InitSubCategoryRouter(router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	subCategoryController, err := wire.InitSubCategoryRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo SubCategoryRouterHandler: " + err.Error())
	}

	router.GET("", controller.Build(subCategoryController.GetAll, logger))
	router.GET("/name/:name", controller.Build(subCategoryController.GetByName, logger))
	router.GET("/:id", controller.Build(subCategoryController.GetByID, logger))
	router.POST("", controller.Build(subCategoryController.Create, logger))
	router.PUT("/:id", controller.Build(subCategoryController.Update_Put, logger))
	router.DELETE("/:id", controller.Build(subCategoryController.Delete, logger))
}
