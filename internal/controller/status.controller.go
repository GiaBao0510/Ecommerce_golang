package controller

import (
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

type StatusController struct {
	statusService service.IStatusService
}

// hàm khởi tạo
func NewStatusController(statusService service.IStatusService) *StatusController {
	return &StatusController{
		statusService: statusService,
	}
}

// -------- Xử lý các request liên quan đến CRUD của Status ở đây ----------
// GET /statuses/:id
func (s *StatusController) GetStatusByID(c *gin.Context) {

	id := c.Param("id")       // Lấy ID từ URL
	id_int := verifyID(c, id) // Validate ID và chuyển đổi sang int32

	result, err := s.statusService.GetStatusByID(c, id_int)
	if err != nil {
		response.ErrorResponse(c, response.StatusBadRequest,
			"Failed to get status by ID: "+err.Error())
		return
	}

	response.SuccessResponse(c, response.StatusOK, result)
}

// GET /statuses
func (s *StatusController) GetAllStatuses(c *gin.Context) {
	result, err := s.statusService.GetAllStatuses(c)
	if err != nil {
		response.ErrorResponse(c, response.StatusBadRequest,
			"Failed to get all statuses: "+err.Error())
		return
	}

	response.SuccessResponse(c, response.StatusOK, result)
}

// POST /statuses
func (s *StatusController) CreateStatus(c *gin.Context) {

	input := models.Status{}

	//Parse JSON body vào struct Status
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.StatusBadRequest,
			"Invalid input data: "+err.Error())
		return
	}

	// Gọi service để tạo mới status
	result, err := s.statusService.CreateStatus(c, &input)
	if err != nil {
		response.ErrorResponse(c, response.StatusBadRequest,
			"Failed to create status: "+err.Error())
		return
	}

	response.SuccessResponse(c, response.StatusOK, gin.H{"id": result})
}

// PUT /statuses/:id
func (s *StatusController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")       // Lấy ID từ URL
	id_int := verifyID(c, id) // Validate ID và chuyển đổi sang int32

	input := models.Status{}

	//Parse JSON body vào struct Status
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}

	// Gọi service để cập nhật status
	if err := s.statusService.UpdateStatus(c, int32(id_int), &input);err != nil {
		response.ErrorResponse(c, response.StatusBadRequest, "Failed to update status: "+err.Error())
		return
	}
	response.SuccessResponse(c, response.StatusOK, "Status updated successfully")
}

// DELETE /statuses/:id
func (s *StatusController) DeleteStatus(c *gin.Context) {
	id := c.Param("id")       // Lấy ID từ URL
	id_int := verifyID(c, id) // Validate ID và chuyển đổi sang int32


	// Gọi service để xóa status
	if err := s.statusService.DeleteStatus(c, int32(id_int)); err != nil {
		response.ErrorResponse(c, response.StatusBadRequest, "Failed to delete status: "+err.Error())
		return
	}
	response.SuccessResponse(c, response.StatusOK, "Status deleted successfully")
}

// ------------ Validate dữ liệu đầu vào cho Status ------------
func verifyID(c *gin.Context, id string) int32 {
	id_int, err := strconv.Atoi(id)

	// Kiểm tra nếu có lỗi khi chuyển đổi ID từ string sang int
	if err != nil {
		response.ErrorResponse(c, response.StatusBadRequest, "Mã ID không phải số nguyên hợp lệ")
		return 0
	}

	if id_int <= 0 {
		response.ErrorResponse(c, response.StatusBadRequest, "Mã ID phải lớn hơn 0")
		return 0
	}

	return int32(id_int)
}

func verifyStatusName(c *gin.Context, name string) {
	if name == "" {
		response.ErrorResponse(c, response.StatusBadRequest, "Status name cannot be empty")
	}
}
