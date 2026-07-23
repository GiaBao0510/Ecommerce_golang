package controller

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RolePermissionController struct {
	rp     service.IRolePermissionService
	logger *zap.Logger
}

func NewRolePermissionController(rp service.IRolePermissionService, logger *zap.Logger) *RolePermissionController {
	return &RolePermissionController{rp: rp, logger: logger}
}

// -------- Xử lý các request liên quan đến CRUD của RolePermission ở đây ----------
// GET /role-permissions/role/:id
func (c *RolePermissionController) GetPermissionsByRoleID(ctx *gin.Context) error {
	// Lấy ID từ param
	id := ctx.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	result, err := c.rp.GetPermissionsByRoleID(ctx, id_int)
	if err != nil {
		return err
	}

	response.Success_Response(ctx, 200, "Permissions retrieved successfully", result)
	return nil
}

// GET /role-permissions/permission/:id
func (c *RolePermissionController) GetRolesByPermissionID(ctx *gin.Context) error {
	// Lấy ID từ param
	id := ctx.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	result, err := c.rp.GetRolesByPermissionID(ctx, id_int)
	if err != nil {
		return err
	}
	
	response.Success_Response(ctx, 200, "Roles retrieved successfully", result)
	return nil
}

// POST /role-permissions
func (c *RolePermissionController) Create(ctx *gin.Context) error {
	// Lấy dữ liệu từ body
	input := models.Role_Permission{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	result, err := c.rp.Create(ctx, &input)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, 201, "Role-Permission created successfully", result)
	return nil
}

// PUT /role-permissions/:role_id/:permission_id
func (c *RolePermissionController) Update_Put(ctx *gin.Context) error { 
	
	// Lấy dữ liệu đầu vào
	input := models.Role_Permission{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	if err := c.rp.Update_Put(ctx, &input); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, 200, "Role-Permission updated successfully", nil)
	return nil
}	

// DELETE /role-permissions/:role_id/:permission_id
func (c *RolePermissionController) Delete(ctx *gin.Context) error {

	// Lấy dữ liệu đầu vào
	input := models.Role_Permission{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	if err := c.rp.Delete(ctx, &input); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(ctx, 200, "Role-Permission deleted successfully", nil)
	return nil
}