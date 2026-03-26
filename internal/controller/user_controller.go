package controller

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

//hàm khởi tạo
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (obj *UserController) GetInforController( c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": obj.userService.UserRepo.GetInfo(),
	})
}