package initialize

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers"
	"github.com/gin-gonic/gin"
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
		r = gin.Default()          // Sử dụng gin.Default() để tự động thêm các middleware như Logger và Recovery, giúp ghi log chi tiết và phục hồi sau lỗi

		// Ngược lại, nếu ở môi trường sản xuất, tắt chế độ debug để tăng hiệu suất và bảo mật
	} else {
		gin.SetMode(gin.ReleaseMode) // Tắt chế độ debug để giảm log chi tiết và tăng hiệu suất
	}

	r = gin.New()                // Sử dụng gin.New() để tạo một instance Gin mới mà không có middleware mặc định, giúp giảm overhead và tăng hiệu suất trong môi trường sản xuất


	// ---Bước 2: Đăng ký middleware chung------------------------------------------------------
	// Middleware chung cho tất cả các route, ví dụ như CORS, Logging, Recovery, v.v.
	// ------------------------------------------------------------------------------------------
	r.Use(
		// Loggin
		// Limiter
		gin.Recovery(), // Middleware Recovery giúp phục hồi sau panic, tránh crash server và ghi log lỗi
		// cors
		// Authen
		// Permission
	) 

	// ---Bước 3: Định nghĩa route cho từng vai trò------------------------------------------------------
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/api")
	{
		MainGroup.GET("/checkStatus")
	}

	// Định nghĩa route cho từng vai trò
	UserGroup := MainGroup.Group("/user")
	ManagerGroup := MainGroup.Group("/manager")

	userRouter.InitProductRouter(UserGroup)
	userRouter.InitUserRouter(UserGroup)

	managerRouter.InitAdminRouter(ManagerGroup.Group("/admin"))
	managerRouter.InitUserRouter(ManagerGroup.Group("/user"))
	managerRouter.InitStatusRouter(
		ManagerGroup.Group("/status"),
		global.DB,
		global.Logger.Error)
	managerRouter.InitRolesRouter(
		ManagerGroup.Group("/roles"), 
		global.DB, 
		global.Logger.Error)

	return r
}
