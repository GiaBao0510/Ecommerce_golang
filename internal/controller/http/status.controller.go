package http

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StatusController struct {
	statusService service.IStatusService
	logger *zap.Logger
}

// hàm khởi tạo
func NewStatusController(statusService service.IStatusService, logger *zap.Logger) *StatusController {
	return &StatusController{ statusService: statusService, logger: logger,	}
}

// Build nhận logger từ controller instance
func (s *StatusController) Build(handeler AppHandler) gin.HandlerFunc {
	return Build(handeler, s.logger)
}

// -------- Xử lý các request liên quan đến CRUD của Status ở đây ----------
// GET /statuses/:id
func (s *StatusController) GetStatusByID(c *gin.Context) error {

	var param dto.ID_Param
	if err := c.ShouldBindUri(&param); err !=nil {
		return HandleValidationError(err)
	}

	result, err := s.statusService.GetByID(c, param.ID)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "Status retrieved successfully", result)
	return nil
}

// GET /statuses
func (s *StatusController) GetAllStatuses(c *gin.Context) error {
	result, err := s.statusService.GetAll(c)
	if err != nil {
		return err
	}

	response.Success_Response(c, http.StatusOK, "All statuses retrieved successfully", result)
	return nil
}

// POST /statuses
func (s *StatusController) CreateStatus(c *gin.Context) error {

	input := models.Status{}

	//Parse JSON body vào struct Status
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}

	// Gọi service để tạo mới status
	result, err := s.statusService.Create(c, &input)
	if err != nil {
		return apperrors.NewBadRequestError("Failed to create status: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Status created successfully", gin.H{"id": result})
	return nil
}

// PUT /statuses/:id
func (s *StatusController) Update_Put(c *gin.Context) error {
	
	var param dto.ID_Param 
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}


	input := models.Status{}

	//Parse JSON body vào struct Status
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}

	// Gọi service để cập nhật status
	if err := s.statusService.Update_Put(c, param.ID, &input); err != nil {
		return apperrors.NewBadRequestError("Failed to update status: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Status updated successfully", nil)
	return nil
}

// PUT /statuses/:id
func (s *StatusController) Update_Patch(c *gin.Context) error {
	
	var param dto.ID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	input := models.Status{}

	//Parse JSON body vào struct Status
	if err := c.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}

	// Gọi service để cập nhật status
	if err := s.statusService.Update_Patch(c, param.ID, &input); err != nil {
		return apperrors.NewBadRequestError("Failed to update status: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Status updated successfully", nil)
	return nil
}

// DELETE /statuses/:id
func (s *StatusController) DeleteStatus(c *gin.Context) error {
	
	var param dto.ID_Param
	if err := c.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}

	// Gọi service để xóa status
	if err := s.statusService.Delete(c, param.ID); err != nil {
		return apperrors.NewBadRequestError("Failed to delete status: " + err.Error())
	}

	response.Success_Response(c, http.StatusOK, "Status deleted successfully", nil)
	return nil
}
