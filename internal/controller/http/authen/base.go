package authen

import "github.com/gin-gonic/gin"

type AuthenController struct {
	loginCtrl    *LoginController
	logoutCtrl   *LogoutController
	registerCtrl *RegisterController
}

// Hàm khởi tạo
func NewAuthenController(
	loginCtrl *LoginController,
	logoutCtrl *LogoutController,
	registerCtrl *RegisterController,
) IAuthenController {
	return &AuthenController{
		loginCtrl:    loginCtrl,
		logoutCtrl:   logoutCtrl,
		registerCtrl: registerCtrl,
	}
}

func (ctr *AuthenController) Login(ctx *gin.Context) error {
	return ctr.loginCtrl.Login(ctx)
}

func (ctr *AuthenController) Logout(ctx *gin.Context) error {
	return ctr.logoutCtrl.Logout(ctx)
}

func (ctr *AuthenController) Register(ctx *gin.Context) error {
	return ctr.registerCtrl.Register(ctx)
}
