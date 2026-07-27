package controller

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService service.IUserService
	logger *zap.Logger
}

// hàm khởi tạo
func NewUserController(userService service.IUserService, logger *zap.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger: logger,
	}
}

// CRUD
func(ctr *UserController)	GetByID(c *gin.Context) error {
	id := c.Param("id") // Lấy ID từ URL
	result, err := ctr.
}

func(ctr *UserController)	Create(c *gin.Context) error
func(ctr *UserController)	Update_Put(c *gin.Context) error
func(ctr *UserController)	Update_Patch(c *gin.Context) error
func(ctr *UserController)	Delete(c *gin.Context) error

	// List operations 
func(ctr *UserController)	GetAll(c *gin.Context) error

	// Search operations
func(ctr *UserController)	GetUserByEmail(c *gin.Contextg) error
func(ctr *UserController)	GetUserByPhone(c *gin.Context) error

	// Relationship operations

	// Update other operations
func(ctr *UserController)	UpdateUserAvatar_PATCH(c *gin.Context) error

