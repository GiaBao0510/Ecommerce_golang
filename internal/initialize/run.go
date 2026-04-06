package initialize

import (
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()

	
	InitLogger()


	InitPostgreSQL()
	
	InitRedis()

	// Hiển thị cấu hình đã load để kiểm tra
	DisplayConfig()

	r := InitRouter()
	r.Run(":8080")
}

// Hàm tạm - chỉ hiển thị cấu hình đã load để kiểm tra
func DisplayConfig(){

	fmt.Println("\t ==== PostgreSQL: ===== ")
	fmt.Println("Loading config postgresql: ", global.Config.PostgreSQL.User)
	fmt.Println("Loading config postgresql: ", global.Config.PostgreSQL.DBName)
	fmt.Println("Loading config postgresql: ", global.Config.PostgreSQL.Host)
	fmt.Println("Loading config postgresql: ", global.Config.PostgreSQL.Port)
	fmt.Println("\t ==== Logger: ===== ")
	global.Logger.Info("Init logger successfully", zap.String("ok", "success"))
	fmt.Println("Load Logger file: ", global.Config.Logger.LogFile )
	fmt.Println("Load Logger level: ", global.Config.Logger.Loglevel )
	fmt.Println("Load Logger compress: ", global.Config.Logger.Compress )
	fmt.Println("Load Logger max backups: ", global.Config.Logger.MaxBackups )
	fmt.Println("Load Logger max size: ", global.Config.Logger.MaxSize )
	
}