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

type RolesController struct {
	roleService service.IRolesService
	logger *zap.Logger
}

// hàm khởi tạo
func NewRolesController(roleService service.IRolesService, logger *zap.Logger) *RolesController {
	return &RolesController{ roleService: roleService, logger: logger,	}
}

// -------- Xử lý các request liên quan đến CRUD của Roles ở đây ----------
//GET /roles/:id
func (r *RolesController) GetByID(c *gin.Context) error {
	
	//Lấy Thông số ở Param
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	result, err := r.roleService.GetByID(c, id_int)
	if err != nil {
		return err
	}

	response.Success_Response(c, 200, "Role retrieved successfully", result)
	return nil	
}

//GET /roles
func (r *RolesController) GetAll(c *gin.Context) error {
	result, err := r.roleService.GetAll(c)
	if err != nil {
		return err
	}

	response.Success_Response(c, 200, "All roles retrieved successfully", result)
	return nil
}

//POST /roles
func (r *RolesController) Create(c *gin.Context) error {
	
	// Lấy dữ liệu từ body
	input := models.Role{}
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	// xác minh dữ liệu đầu vào
	if util.VerifyName(input.Role_name) != nil {
		return apperrors.NewBadRequestError("Role name cannot be empty")
	}

	result, err := r.roleService.Create(c, &input)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(c, 200, "Role created successfully", gin.H{"id": result})
	return nil

}

//PUT /roles/:id
func (r *RolesController) Update(c *gin.Context) error {

	//Lấy dữ liệu đầu vào
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	input := models.Role{}
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input: " + err.Error())
	}

	if util.VerifyName(input.Role_name) != nil {
		return apperrors.NewBadRequestError("Role name cannot be empty")
	}

	if err := r.roleService.Update(c, id_int, &input); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(c, 200, "Role updated successfully", nil)
	return nil
}

//DELETE /roles/:id
func (r *RolesController) Delete(c *gin.Context) error {
	id := c.Param("id")
	id_int, err := util.VerifyID(id)
	if err != nil {
		return apperrors.NewBadRequestError("Invalid ID: " + err.Error())
	}

	if err := r.roleService.Delete(c,id_int); err != nil {
		return apperrors.NewInternalServerError(err)
	}

	response.Success_Response(c, 200, "Role deleted successfully", nil)
	return nil
}