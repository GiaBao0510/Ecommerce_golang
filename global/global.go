package global

import (
	"github.com/GiaBao0510/Ecommerce_golang/pkg/logger"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	PostgreSQL *gorm.DB
)