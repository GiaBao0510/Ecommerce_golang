package initialize

import "github.com/GiaBao0510/Ecommerce_golang/internal/middleware"

// Tại đây nó sẽ khởi chạy các tác vụ chạy ngầm
func InitializeBackgroundTasks() {
	go middleware.CleanUpClients() // Dọn dẹp các limiter của các IP đã lâu không gửi request
}