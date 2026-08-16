package authen

import "github.com/gin-gonic/gin"

type LogoutController struct{}

func NewLogoutController() *LogoutController {
	return &LogoutController{}
}

func (L *LogoutController) Logout(ctx *gin.Context) error {
	return nil
}