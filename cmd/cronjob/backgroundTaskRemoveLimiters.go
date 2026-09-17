package cronjob

import "github.com/GiaBao0510/Ecommerce_golang/internal/middleware"

// Goroutine chạy ngầm để dọn dẹp các limiter của các IP đã lâu không gửi request
func InitializeBackgroundTasks() {
	go middleware.CleanUpClients() // Dọn dẹp các limiter của các IP đã lâu không gửi request
}