package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BrandRouter struct{}

func (r *BrandRouter) InitBrandRouter(router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	brandController, err := wire.InitBrandRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo BrandRouterHandler: " + err.Error())
	}

	router.GET("", controller.Build(brandController.GetAll, logger))
	router.GET("/name/:name", controller.Build(brandController.GetByName, logger))
	router.GET("/:id", controller.Build(brandController.GetByID, logger))
	router.POST("", controller.Build(brandController.Create, logger))
	router.PUT("/:id", controller.Build(brandController.Update_Put, logger))
	router.DELETE("/:id", controller.Build(brandController.Delete, logger))
}
