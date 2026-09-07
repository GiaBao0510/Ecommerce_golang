package authen

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/GiaBao0510/Ecommerce_golang/global"
	//controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	//"github.com/GiaBao0510/Ecommerce_golang/internal/models"
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
	
	response.Success_Response(ctx, http.StatusOK, "Refresh token successfully", nil)
	return nil
}