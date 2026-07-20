package routers

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetUpRouter khởi tạo gin Engine và đăng ký các middleware + route
// Điều này giúp các middleware có thể dùng chung một instace logger duy nhất
func SetUpRouter(logger *zap.Logger) *gin.Engine {

	var HealthController = controller.NewHealthController()

	// Điều này giúp chúng ta có thể kiểm soát middleware.
	// Mặc định (gin.Default()) sẽ có 2 middleware: Logger + Recovery.
	// Nhưng chúng ta muốn custom middleware riêng, nên dùng gin.New() để tạo engine trống.
	r := gin.New()

	/*	--------------------------------------------
		Đăng ký các middleware theo thứ tự rất quan trọng

		Thứ tự middleware như sau:
			1. TraceIDMiddleware → phải chạy ĐẦU TIÊN để set trace_id cho các middleware sau
			2. HTTPLoggerMiddleware → phải sau TraceID để có trace_id trong log
			3. Recovery → bắt panic để app không crash
			4. AuthenMiddleware → xác thực JWT
		-------------------------------------------*/
	r.Use(
		middleware.RecoveryMiddleware(),         // Bắt panic, trả 500 thay vì crash
		middleware.TraceID_Middleware(),         // Sinh ra trace_id cho các requets
		middleware.HttpLoggerMiddleware(logger), // Ghi log cho mọi request (method, path, status, latency, trace_id)

		//HSTS(),
		//HttpRedirection(),
		middleware.AuthenMiddleware(),
	)

	v1 := r.Group("/v1/api")
	{
		v1.GET("/health", HealthController.CheckHealth)
		v1.GET("/health/live", HealthController.Live)
	}

	// Ghi log thông tin khi router được khởi tạo xong
	logger.Info("Router initialized successfully", zap.String("base_path", "/v1/api"))
	return r
}
