package initialize

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/middleware"
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	// ---Bước 1: Khởi tạo router------------------------------------------------------------------------------------
	// Nếu dự án ở môi trường phát triển "dev", có thể bật chế độ debug để dễ dàng phát hiện lỗi
	// Nếu dự án ở môi trường sản xuất "prod", tắt chế độ debug để tăng hiệu suất và bảo mật
	// ------------------------------------------------------------------------------------------
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode) // Bật chế độ debug để hiển thị log chi tiết và thông tin lỗi
		gin.ForceConsoleColor()    // Buộc hiển thị màu sắc trong log để dễ phân biệt các loại log khác nhau (thông tin, cảnh báo, lỗi)
		//r = gin.Default()          // Sử dụng gin.Default() để tự động thêm các middleware như Logger và Recovery, giúp ghi log chi tiết và phục hồi sau lỗi

		// Ngược lại, nếu ở môi trường sản xuất, tắt chế độ debug để tăng hiệu suất và bảo mật
	} else {
		gin.SetMode(gin.ReleaseMode) // Tắt chế độ debug để giảm log chi tiết và tăng hiệu suất
	}

	r = gin.New() // Sử dụng gin.New() để tạo một instance Gin mới mà không có middleware mặc định, giúp giảm overhead và tăng hiệu suất trong môi trường sản xuất

	/* ---Bước 2: Đăng ký middleware chung------------------------------------------------------
		THỨ TỰ ĐĂNG KÝ RẤT QUAN TRỌNG — middleware chạy TRƯỚC bọc quanh middleware
	chạy SAU (giống như vỏ hành). Lý do thứ tự dưới đây:
	
	1. RealIP     : không phụ thuộc middleware nào khác, cần chạy sớm vì
	Logger cần đọc RealIPKey từ context.

	2. TraceID: Sinh UUID duy nhất cho mỗi request
	→ Lưu vào context để tất cả middleware/handler SAU có thể dùng
	→ Nếu đặt sau Logger, Logger sẽ không có trace_id để ghi

	3. Traing: span nên bọc quanh CÀNG NHIỀU middleware phía sau
	càng tốt để đo được toàn bộ thời gian xử lý.

	4.  Recovery — Bắt panic, ngăn app crash
	→ Nếu handler nào bị panic (lỗi không mong đợi), Recovery bắt lại
	→ Trả về 500 thay vì để server tắt
	→ PHẢI đăng ký để đảm bảo an toàn cho production

	5.  Logger — ghi access log — đặt sau Recovery để nếu có panic
	→ Ghi access log: method, path, status, latency, trace_id
	→ Ghi vào storages/logs/access.log
	→ Dùng Access logger (không phải Error logger)

	6. tương tự Logger, cần status code cuối cùng.

	
	------------------------------------------------------------------------------------------ */
	r.Use(
		middleware.RealIPMiddleware(),
		middleware.TraceID_Middleware(),
		middleware.TracingMiddleware(),
		middleware.RecoveryMiddleware(),
		middleware.HttpLoggerMiddleware(global.Logger.Access),
		middleware.MetricsMiddleware(),

		// [4] CORS — Cho phép cross-origin requests (frontend khác domain)
		// → Chưa implement, sẽ thêm sau
		// middleware.CorsMiddleware(),

		// [5] Rate Limiter — Giới hạn số request từ 1 IP
		// → Ngăn DDoS, brute force
		// → Chưa implement, sẽ thêm sau
		// middleware.RateLimitMiddleware(),

		// [6] Authentication — Xác thực JWT token
		// → Kiểm tra Authorization header
		// → PHẢI sau Recovery (để Recovery bắt được panic nếu auth bị lỗi)
		// → Chưa implement đầy đủ, bật khi sẵn sàng
		// middleware.AuthenMiddleware(),
	)

	/* ==================================================
	// Metric Endpoint — /metrics
	Middleware MetricsMiddleware() đã thu thập dữ liệu vào registry.
	// Thêm route "/metrics" để Prometheus server có endpoint để scrape định kỳ.
	// Đặt ngoài MainGroup (không tiền tố /v1/api) theo convention chuẩn của
	// Prometheus; không cần qua Auth/RateLimit vì đây là traffic nội bộ
	// (Prometheus server), nên whitelist IP nội bộ ở tầng network/firewall.
	==================================================== */
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// ================================================================
	// ROUTES
	// ================================================================
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/api")
	{
		MainGroup.GET("/checkStatus")
	}

	// Định nghĩa route cho từng vai trò
	UserGroup := MainGroup.Group("/user")
	ManagerGroup := MainGroup.Group("/manager")

	// Định nghĩa route cho vai trò user
	userRouter.InitProductRouter(UserGroup)
	userRouter.InitUserRouter(
		UserGroup.Group("/user"),
		global.DB,
		global.Logger.Error,
	)

	// Định nghĩa route cho vai trò admin
	managerRouter.InitAdminRouter(ManagerGroup.Group("/admin"))
	managerRouter.InitUserRouter(
		ManagerGroup.Group("/user"),
		global.DB,
		global.Logger.Error,
	)
	managerRouter.InitStatusRouter(
		ManagerGroup.Group("/status"),
		global.DB,
		global.Logger.Error)
	managerRouter.InitRolesRouter(
		ManagerGroup.Group("/roles"),
		global.DB,
		global.Logger.Error)
	managerRouter.InitPermissionRouter(
		ManagerGroup.Group("/permission"),
		global.DB,
		global.Logger.Error)
	managerRouter.InitRolePermissionRouter(
		ManagerGroup.Group("/role_permission"),
		global.DB,
		global.Logger.Error)

	return r
}
