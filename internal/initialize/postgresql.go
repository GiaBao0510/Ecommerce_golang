package initialize

import (
	"fmt"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Tạo hàm kiểm tra lỗi. Nếu xảy ra lỗi thì nó sẽ thông báo chi tiết nơi xảy và và đồng thời lưu trong log
func CheckErrorPanic(err error, message string) {
	if err != nil {
		global.Logger.Error(message, zap.Error(err))
		panic(fmt.Sprintf("%s: %v", message, err)) // Panic để dừng chương trình nếu có lỗi nghiêm trọng, đồng thời cung cấp thông tin chi tiết về lỗi
	}
}

// Các hàm khởi tạo cho PostgreSQL sẽ được đặt ở đây
func InitPostgreSQL() {
	m := global.Config.PostgreSQL

	// Tạo chuỗi kết nối (DSN) cho PostgreSQL
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	var s = fmt.Sprintf(dsn, m.Host, m.User, m.Password, m.DBName, m.Port)

	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,                               // Tắt giao dịch mặc định để cải thiện hiệu suất, nhưng cần cẩn thận khi sử dụng
		Logger:                 logger.Default.LogMode(logger.Warn), // Thiết lập logger của GORM để chỉ log các cảnh báo và lỗi, giúp giảm bớt log không cần thiết trong quá trình phát triển
	})
	CheckErrorPanic(err, "InitPostgreSQL: Failed to connect to PostgreSQL")

	global.Logger.Info("InitPostgreSQL: Successfully connected to PostgreSQL")
	global.PostgreSQL = db

	// Set pool
	SetPool()

	// Migrate tables
	migrateTables()

}

// Hàm này sẽ thiết lập pool kết nối cho PostgreSQL
// Việc này thiết lập mở nhóm kết nối tối đa, số lượng kết nối nhàn rỗi tối đa và thời gian sống tối đa của kết nối
func SetPool() {
	p := global.Config.PostgreSQL
	sqlDB, err := global.PostgreSQL.DB()
	if err != nil {
		global.Logger.Error("SetPool: Failed to get database from GORM", zap.Error(err))
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(p.MaxIdleConns))    // Thiết lập thời gian tối đa của một kết nối nhàn rỗi trước khi bị đóng
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)                      // thiết lập giới hạn số lượng kết nối tối đa để tránh quá tải cơ sở dữ liệu
	sqlDB.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetime)) // SAu khi kết nối tồn tại hơn thời gian được thiết lập ở đây thì nó sẽ bị đóng vào loại khỏi pool
}

// Hàm này sẽ chạy các migration để tạo bảng nếu chưa tồn tại
func migrateTables() {
	// Tại đây câu lệnh này sẽ tự động tạo bảng dựa trên các struct đã định nghĩa trong package po.
	err := global.PostgreSQL.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("migrateTables: Failed to migrate tables", zap.Error(err))
	}
}
