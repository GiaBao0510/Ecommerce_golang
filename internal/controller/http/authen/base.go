package authen

import "github.com/gin-gonic/gin"

type AuthenController struct {
	loginCtrl        *LoginController
	loginGoogleCtrl  *LoginGoogleController
	logoutCtrl       *LogoutController
	registerCtrl     *RegisterController
	refreshTokenCtrl *RefreshTokenController
}

// Hàm khởi tạo
func NewAuthenController(
	loginCtrl *LoginController,
	loginGoogleCtrl *LoginGoogleController,
	logoutCtrl *LogoutController,
	registerCtrl *RegisterController,
	refreshTokenCtrl *RefreshTokenController,
) IAuthenController {
	return &AuthenController{
		loginCtrl:        loginCtrl,
		loginGoogleCtrl:  loginGoogleCtrl,
		logoutCtrl:       logoutCtrl,
		registerCtrl:     registerCtrl,
		refreshTokenCtrl: refreshTokenCtrl,
	}
}

func (ctr *AuthenController) Login(ctx *gin.Context) error {
	return ctr.loginCtrl.Login(ctx)
}

func (ctr *AuthenController) Login_Google(ctx *gin.Context) error {
	return ctr.loginGoogleCtrl.Login_Google(ctx)
}

func (ctr *AuthenController) Logout(ctx *gin.Context) error {
	return ctr.logoutCtrl.Logout(ctx)
}

func (ctr *AuthenController) Register(ctx *gin.Context) error {
	return ctr.registerCtrl.Register(ctx)
}

func (ctr *AuthenController) RefreshToken(ctx *gin.Context) error {
	return ctr.refreshTokenCtrl.RefreshToken(ctx)
}
