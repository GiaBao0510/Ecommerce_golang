package initialize

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/logger"
)

func InitLogger() {
	// Khởi tạo logger và gán vào biến toàn cục
	// global.Logger kiểu là *logger.LoggerZap — đây là struct wrapper bọc *zap.Logger
	// global.Logger.Logger là *zap.Logger (raw, bên trong wrapper)
	global.Logger = logger.NewLogger(global.Config.Logger) 
}
