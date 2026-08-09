package user

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/wire"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailRouter struct{}

func (r *EmailRouter) InitEmailRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	emailController, err := wire.InitVerifyRouterHandler(db, logger)
	if err != nil {
		panic("Lỗi khi khởi tạo EmailRouterHandler: " + err.Error())
	}

	// Private routes
	Router.POST("/verify", controller.Build(emailController.VerifyEmail, logger))

	// Public routes
	Router.GET("/get_verification_code", controller.Build(emailController.SendVerificationEmail, logger))
	Router.POST("/verify-otp", controller.Build(emailController.VerifyOTP_viaEmail, logger))
}
