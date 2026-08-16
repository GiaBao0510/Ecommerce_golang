package authen

import "github.com/gin-gonic/gin"

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (L *LoginController) Login(ctx *gin.Context) error {
	return nil
}