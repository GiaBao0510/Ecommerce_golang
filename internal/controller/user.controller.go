package controller

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
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
	var param dto.UUID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	result, err := ctr.userService.GetByID(c, param.UUID)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User retrieved successfully", result)
	return nil
}

func(ctr *UserController)	Create(c *gin.Context) error {
	var input models.CreateUsersRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return HandleValidationError(err)
	}

	result, err := ctr.userService.Create(c, input)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusCreated, "User created successfully", result)
	return nil
}
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

