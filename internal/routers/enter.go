package routers

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers/manager"
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers/public"
	"github.com/GiaBao0510/Ecommerce_golang/internal/routers/user"
)

// Ở tệp tin này sẽ tổng hợp lại các router từ user, admin, ...
type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
	Public public.PublicRouterGroup
}

var RouterGroupApp = new(RouterGroup)	// Khởi tạo biến toàn cục này để giúp trỏ về