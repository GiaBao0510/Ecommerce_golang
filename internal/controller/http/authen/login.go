package authen

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
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

	// thiết lập SameSite để giảm nguy cơ tấn công CSRF
	ctx.SetSameSite(http.SameSiteLaxMode)

	// Tạo cookie với token và thiết lập các thuộc tính bảo mật
	ctx.SetCookie(
		"access_token", 
		result.AccessToken, 
		global.Config.Authentication.JWT.AccessTokenExpirationMinutes * 60, 
		"/",
		"",
		false, // Không chỉ gửi cookie qua HTTPS (vì đang phát triển trên localhost, nên đặt là false)
		true,  // Chỉ cho phép cookie được truy cập bởi trình duyệt (không thể truy cập bằng JavaScript)
	)

	ctx.SetCookie(
		"refresh_token", 
		result.RefreshToken, 
		global.Config.Authentication.JWT.RefreshTokenExpirationDays * 60 * 60 * 24, 
		"/",
		"",
		false, // Không chỉ gửi cookie qua HTTPS (vì đang phát triển trên localhost, nên đặt là false)
		true,  // Chỉ cho phép cookie được truy cập bởi trình duyệt (không thể truy cập bằng JavaScript)
	)

	// Gửi phản hồi
	response.Success_Response(ctx, 200, "Login successful", result)
	return nil
}