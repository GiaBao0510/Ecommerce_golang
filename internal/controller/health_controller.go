package controller

import (
	"runtime"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
)

// Structure HealthController để kiểm tra sức khỏe của ứng dụng
type HealthController struct {
	//healthService *service.HealthService
	startTime time.Time
}

// Hàm khởi tạo mới cho HealthController
func NewHealthController () *HealthController {
	return &HealthController{
		startTime: time.Now(),
	}
}

// Hàm kiểm tra sức khỏe của ứng dụng
func (obj *HealthController) CheckHealth(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, nil,)
}


// GET/ health/live - kiểm tra xem server có bị treo không
func (obj *HealthController) Live(c *gin.Context) {
	
	// Kiểm tra xem server có bị treo không bằng cách kiểm tra thời gian hoạt động
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	uptime := time.Since(obj.startTime)

	response.SuccessResponse(
		c, 
		response.ErrorCodeSuccess,
		gin.H{
			"status": "alive",
			"uptime": util.FotmatUptime(uptime),
			"memory": gin.H{
				"used_mb": memStats.Alloc / 1024 / 1024,
				"total_mb": memStats.Sys / 1024 / 1024,
			},
			"goroutines": runtime.NumGoroutine(),
		},
	)
}