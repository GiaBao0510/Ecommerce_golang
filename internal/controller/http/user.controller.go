package http

import (
	"database/sql"
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/user"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService service.IUserService
	logger      *zap.Logger
}
 
// hàm khởi tạo
func NewUserController(userService service.IUserService, logger *zap.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

// CRUD
func (ctr *UserController) GetByID(c *gin.Context) error {
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

func (ctr *UserController) Create(c *gin.Context) error {
	var input models.CreateUsersRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return HandleValidationError(err)
	}

	result, err := ctr.userService.Create(c, &input)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusCreated, "User created successfully", result)
	return nil
}

func (ctr *UserController) Update_Put(c *gin.Context) error {
	var param dto.UUID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	var input models.UpdateUsersPutRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		return HandleValidationError(err)
	}

	if err := ctr.userService.Update_Put(c, param.UUID, &input); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User updated successfully", nil)
	return nil
}

func (ctr *UserController) Update_Patch(c *gin.Context) error {
	var param dto.UUID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	var input models.UpdateUsersPatchRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		return HandleValidationError(err)
	}

	if err := ctr.userService.Update_Patch(c, param.UUID, &input); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User updated successfully", nil)
	return nil
}

func (ctr *UserController) Delete(c *gin.Context) error {
	var param dto.UUID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	if err := ctr.userService.Delete(c, param.UUID); err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User deleted successfully", nil)
	return nil
}

// List operations
func (ctr *UserController) GetAll(c *gin.Context) error {
	result, err := ctr.userService.GetAll(c)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Users retrieved successfully", result)
	return nil
}

// Search operations
func (ctr *UserController) GetUserByEmail(c *gin.Context) error {
	var param dto.Email_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	result, err := ctr.userService.GetUserByEmail(c, param.Email)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User retrieved successfully", result)
	return nil
}

func (ctr *UserController) GetUserByPhone(c *gin.Context) error {
	var param dto.Phone_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	result, err := ctr.userService.GetUserByPhone(c, sql.NullString{String: param.Phone, Valid: true})
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "User retrieved successfully", result)
	return nil
}

// Relationship operations

// Update other operations
// func(ctr *UserController)	UpdateUserAvatar_PATCH(c *gin.Context) error {

// }
