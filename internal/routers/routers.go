package routers

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{ 

	var HealthController = controller.NewHealthController()

	r := gin.New()

	r.GET("/health", HealthController.CheckHealth)
	r.GET("/health/live", HealthController.Live)

	return  r
}