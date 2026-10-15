package manager

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopOwnerRouter struct{}

func (r *ShopOwnerRouter) InitShopOwnerRouter(router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	shopOwnerController, err := wire.InitShopOwnerRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo ShopOwnerRouterHandler: " + err.Error())
	}

	router.GET("", controller.Build(shopOwnerController.GetAll, logger))
	router.GET("/name/:name", controller.Build(shopOwnerController.GetByName, logger))
	router.GET("/:id", controller.Build(shopOwnerController.GetByID, logger))
	router.POST("", controller.Build(shopOwnerController.Create, logger))
	router.PUT("/:id", controller.Build(shopOwnerController.Update_Put, logger))
	router.PATCH("/:id/name", controller.Build(shopOwnerController.UpdateName_Patch, logger))
	router.PATCH("/:id/description", controller.Build(shopOwnerController.UpdateDescription_Patch, logger))
	router.PATCH("/:id/avatar", controller.Build(shopOwnerController.UpdateAvatar_Patch, logger))
	router.DELETE("/:id", controller.Build(shopOwnerController.Delete, logger))
}
