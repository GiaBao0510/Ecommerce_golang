package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductDetailRouter struct{}

func (r *ProductDetailRouter) InitProductDetailRouter(router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	productDetailController, err := wire.InitProductDetailRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo ProductDetailRouterHandler: " + err.Error())
	}

	router.GET("", controller.Build(productDetailController.GetAll, logger))
	router.GET("/:id", controller.Build(productDetailController.GetByID, logger))
	router.POST("", controller.Build(productDetailController.Create, logger))
	router.PUT("/:id", controller.Build(productDetailController.Update_Put, logger))
	router.PATCH("/:id/price", controller.Build(productDetailController.UpdatePrice_Patch, logger))
	router.PATCH("/:id/like-count", controller.Build(productDetailController.UpdateLikeCount_Patch, logger))
	router.PATCH("/:id/quantity", controller.Build(productDetailController.UpdateQuantity_Patch, logger))
	router.DELETE("/:id", controller.Build(productDetailController.Delete, logger))
}
