package initialize

import (
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	InitLogger()
	InitPostgreSQL()
	InitRedis()

	// Hiển thị cấu hình đã load để kiểm tra
	//DisplayConfig()

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
	fmt.Println("Load Logger file: ", global.Config.Logger.LogFile )
	fmt.Println("Load Logger level: ", global.Config.Logger.Loglevel )
	fmt.Println("Load Logger compress: ", global.Config.Logger.Compress )
	fmt.Println("Load Logger max backups: ", global.Config.Logger.MaxBackups )
	fmt.Println("Load Logger max size: ", global.Config.Logger.MaxSize )

	fmt.Println("\t ==== Redis: ===== ")
	fmt.Println("REDIS Addr: ", global.Config.Redis.Address )
	fmt.Println("REDIS Password: ", global.Config.Redis.Password )
	fmt.Println("REDIS DB: ", global.Config.Redis.DB )
	fmt.Println("REDIS ReadTimeout: ", global.Config.Redis.IdleTimeout )
	fmt.Println("REDIS MaxConnLifetime: ", global.Config.Redis.MaxConnLifetime )
	fmt.Println("REDIS WaitTimeout: ", global.Config.Redis.WaitTimeout )
	fmt.Println("REDIS readTimeout: ", global.Config.Redis.ReadTimeout )
	fmt.Println("REDIS writeTimeout: ", global.Config.Redis.WriteTimeout )

	fmt.Println("\t ==== Server: ===== ")
	fmt.Println("Server Port: ", global.Config.Server.Port )
	fmt.Println("Server Host: ", global.Config.Server.Host )
	fmt.Println("Server Mode: ", global.Config.Server.Mode )
}