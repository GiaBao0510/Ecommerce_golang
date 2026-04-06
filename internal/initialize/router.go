package initialize

import (
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("Befor - Logger")
		c.Next()
		fmt.Println("After - Logger")
	}
}

func ExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("Befor - ExceptionHandler")
		c.Next()
		fmt.Println("After - ExceptionHandler")
	}
}

func HSTS() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("Befor - HSTS")
		c.Next()
		fmt.Println("After - HSTS")
	}
}

func HttpRedirection() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("Befor - HttpRedirection")
		c.Next()
		fmt.Println("After - HttpRedirection")
	}
}

func InitRouter() *gin.Engine{ 

	var HealthController = controller.NewHealthController()

	r := gin.Default()

	r.Use(ExceptionHandler(), HSTS(), HttpRedirection(), Logger(), middleware.AuthenMiddleware())

	v1:= r.Group("/v1/api")
	{
		v1.GET("/health", HealthController.CheckHealth)
		v1.GET("/health/live", HealthController.Live)
	}

	return  r
}