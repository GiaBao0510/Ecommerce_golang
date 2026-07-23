package controller

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

// hàm khởi tạo
func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Hàm đăng ký
func (obj *UserController) RegisterController(c *gin.Context) {
	result := obj.userService.Register("", "")
	
	response.Success_Response(c, http.StatusOK, "Đăng ký thành công", result)
}

// func (obj *UserController) GetInforController( c *gin.Context) {
// 	response.SuccessResponse(c, response.ErrorCodeSuccess, obj.userService.GetInfoService()) 
// }

