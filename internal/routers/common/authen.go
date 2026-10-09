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

func (a *AuthenRouter) InitAuthenRouter(
	Router *gin.RouterGroup, 
	db *sql.DB, 
	queries *database.Queries, 
	logger *zap.Logger,
	authMiddleware gin.HandlerFunc,
) {

	authController, err := wire.InitAuthenRouterHandler(db, queries, logger)
	mailController, err := wire.InitMailRouterHandler(queries, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo ")
	}

	// public routes for authentication
	Router.POST("/register", controller.Build(authController.Register, logger))
	Router.POST("/login", controller.Build(authController.Login, logger))
	Router.POST("/refresh", controller.Build(authController.RefreshToken, logger))

	Router.POST("/login/google", controller.Build(authController.Login_Google, logger))
	Router.GET("/google/callback", controller.Build(authController.Login_Google_Callback, logger))
	Router.GET("/verify-email-mailjet/", controller.Build(mailController.VerifyEmailMailjet, logger))
	Router.GET("/verify-email-mailtrap/", controller.Build(mailController.VerifyEmailMailjet, logger))

	// private routes for authentication (require authentication)
	Router.POST(
		"/logout",
		authMiddleware,
		controller.Build(authController.Logout, logger),
	)

}
