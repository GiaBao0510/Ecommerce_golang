package authen

import (
	"github.com/gin-gonic/gin"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
)

type LoginController struct{
	svc service.IAuthService
}

func NewLoginController(svc service.IAuthService) *LoginController {
	return &LoginController{svc: svc}
}

func (L *LoginController) Login(ctx *gin.Context) error {
	
	input := models.LoginRequest{}

	// Parse JSON body vào struct LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return controller.HandleValidationError(err)
	}

	// Gọi service để thực hiện đăng nhập
	result, err := L.svc.Login(ctx, &input)
	if err != nil {
		return err 
	}

	response.Success_Response(ctx, 200, "Login successful", result)
	return nil
}