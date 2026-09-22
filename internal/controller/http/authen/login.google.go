package authen

import (
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/oauth2"
	"github.com/gin-gonic/gin"
)

type LoginGoogleController struct {
	svc service.IOAuth2Service
}

func NewLoginGoogleController(svc service.IOAuth2Service) *LoginGoogleController {
	return &LoginGoogleController{svc: svc}
}

func (L *LoginGoogleController) Login_Google(ctx *gin.Context) error {
	
	
	return nil
}