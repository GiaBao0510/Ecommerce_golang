package middleware

import (
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: global.Config.Cors.Allowed_origins,
		AllowMethods: global.Config.Cors.Allowed_methods,
		AllowHeaders: global.Config.Cors.Allowed_headers,
		AllowCredentials: global.Config.Cors.Allow_credentials,
		MaxAge: time.Duration(global.Config.Cors.Max_age),
	})
}