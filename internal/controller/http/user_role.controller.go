package http

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/user"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRoleController struct {
	ur     service.IUserRoleService
	logger *zap.Logger
}

func NewUserRoleController(ur service.IUserRoleService, logger *zap.Logger) *UserRoleController {
	return &UserRoleController{ur: ur, logger: logger}
}

// Xử lý các request liên quan đến CRUD của UserRole ở đây
func (c *UserRoleController) Create(ctx *gin.Context) error {
	input := models.UserRole{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	result, err := c.ur.Create(ctx, &input)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, http.StatusCreated, "User role created successfully", result)
	return nil
}

func (c *UserRoleController) Update(ctx *gin.Context) error {
	input := models.UserRole{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	if err := c.ur.Update(ctx, input.Uuid, &input); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, http.StatusOK, "User role updated successfully", nil)
	return nil
}

func (c *UserRoleController) Delete(ctx *gin.Context) error {
	uuid := ctx.Query("uuid")
	roleIDStr := ctx.Query("role_id")

	roleID, err := util.ValidationPositiveInt("Role ID", roleIDStr)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid Role ID: " + err.Error())
	}

	err = c.ur.Delete(ctx, uuid, int32(roleID))
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, http.StatusOK, "User role deleted successfully", nil)
	return nil
}

func (c *UserRoleController) GetUserByRoleID(ctx *gin.Context) error {
	id := ctx.Param("id")
	id_int, err := util.ValidationPositiveInt("ID", id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	result, err := c.ur.GetUserByRoleID(ctx, int32(id_int))
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, http.StatusOK, "Users retrieved successfully", result)
	return nil
}

func (c *UserRoleController) GetRolesByUserID(ctx *gin.Context) error {
	
	uuid := ctx.Param("uuid")
	if err := util.ValidationUUID("UUID", uuid); err != nil {
		return apperrors.NewBadRequestError("Invalid UUID: " + err.Error())
	}

	result, err := c.ur.GetRolesByUserID(ctx, uuid)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, http.StatusOK, "Roles retrieved successfully", result)
	return nil
}
