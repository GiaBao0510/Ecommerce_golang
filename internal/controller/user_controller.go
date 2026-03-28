package controller

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
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
	response.SuccessResponse(c, response.ErrorCodeSuccess, obj.userService.GetInfoService()) 
}