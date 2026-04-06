package initialize

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger) 
}
