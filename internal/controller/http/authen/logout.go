package authen

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

type LogoutController struct{
	svc service.IAuthService
}

func NewLogoutController(svc service.IAuthService) *LogoutController {
	return &LogoutController{svc: svc}
}

func (L *LogoutController) Logout(ctx *gin.Context) error {
	input := models.LogoutRequest{}

	// Parse JSON body vào struct LogoutRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return controller.HandleValidationError(err)
	}

	if err := L.svc.Logout(ctx, &input); err != nil {
		return err
	}

	response.Success_Response(ctx, 200, "Logout successful", nil)
	return nil 
} 