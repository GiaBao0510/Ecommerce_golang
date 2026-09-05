package common

import (
	"database/sql"

	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthenRouter struct{}

func (a *AuthenRouter) InitAuthenRouter(Router *gin.RouterGroup, db *sql.DB, queries *database.Queries, logger *zap.Logger) {
	
	authController, err := wire.InitAuthenRouterHandler(db ,queries ,logger)
	if err != nil {
		panic("Lỗi khi khởi tạo ")
	}

	// public routes for authentication
	Router.POST("/register", controller.Build(authController.Register, logger))
	Router.POST("/login", controller.Build(authController.Login, logger))
	
	// private routes for authentication (require authentication)
	Router.POST("/logout", controller.Build(authController.Logout, logger))
} 