package authen

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
)


type RefreshTokenController struct {
	svc service.IAuthService
}

func NewRefreshTokenController(svc service.IAuthService ) *RefreshTokenController{
	return &RefreshTokenController{svc: svc}
}

func (r *RefreshTokenController) RefreshToken(ctx *gin.Context) error {
	
	input := dto.Token{}

	// Parse JSON body vào struct Token
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return controller.HandleValidationError(err)
	}

	result, er := r.svc.RefreshToken(ctx, &input)
	if er != nil {
		return er
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

	response.Success_Response(ctx, http.StatusOK, "Refresh token successfully", result)
	return nil
}