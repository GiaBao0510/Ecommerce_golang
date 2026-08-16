package authen

import "github.com/gin-gonic/gin"

type IAuthenController interface {
	Login(ctx *gin.Context) error
	Logout(ctx *gin.Context) error
	Register(ctx *gin.Context) error
}