package controller

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PermissionController struct {
	permissionSvc service.IPermissionService
	logger        *zap.Logger
}

// hàm khởi tạo
func NewPermissionController(permissionSvc service.IPermissionService, logger *zap.Logger) *PermissionController {
	return &PermissionController{permissionSvc: permissionSvc, logger: logger}
}

// Build nhận logger từ controller instance
func (ctr *PermissionController) Build(handler AppHandler) gin.HandlerFunc {
	return Build(handler, ctr.logger)
}

// -------- Xử lý các request liên quan đến CRUD của Permission ở đây ----------
// GET /permissions/:id
func (ctr *PermissionController) GetByID(c *gin.Context) error {
	id := c.Param("id")
	id_int, err := util.VerifyID(id)

	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	result, err := ctr.permissionSvc.GetByID(c, id_int)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Permission retrieved successfully", result)
	return nil
}

// GET /permissions
func (ctr *PermissionController) GetAll(c *gin.Context) error {
	result, err := ctr.permissionSvc.GetAll(c)
	if err != nil {
		return err
	}
	response.Success_Response(c, http.StatusOK, "All permissions retrieved successfully", result)
	return nil
}

// POST /permissions
func (ctr *PermissionController) Create(c *gin.Context) error {
	input := models.Permission{}

	// Bind dữ liệu từ request body vào struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	// Thực hiện tạo permission mới thông qua service
	result, err := ctr.permissionSvc.Create(c, &input)
	if err != nil {
		return apperrors.NewBadRequestError("Failed to create permission: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Permission created successfully", result)
	return nil

}

// PUT /permissions/:id
func (ctr *PermissionController) Update_Put(c *gin.Context) error {
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	input := models.Permission{}
	// Bind dữ liệu từ request body vào struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	// Thực hiện cập nhật permission thông qua service
	if err := ctr.permissionSvc.Update_Put(c,id_int, &input); err != nil {
		return apperrors.NewBadRequestError("Failed to update permission: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Permission updated successfully", nil)
	return nil
}

// PATCH /permissions/:id
func (ctr *PermissionController) Update_Patch(c *gin.Context) error {
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	input := models.Permission{}
	// Bind dữ liệu từ request body vào struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	// Thực hiện cập nhật permission thông qua service
	if err := ctr.permissionSvc.Update_Patch(c,id_int, &input); err != nil {
		return apperrors.NewBadRequestError("Failed to update permission: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Permission updated successfully", nil)
	return nil
}

// DELETE /permissions/:id
func (ctr *PermissionController) Delete(c *gin.Context) error {
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	// Thực hiện xóa permission thông qua service
	if err := ctr.permissionSvc.Delete(c, id_int); err != nil {
		return apperrors.NewBadRequestError("Failed to delete permission: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Permission deleted successfully", nil)
	return nil
}
