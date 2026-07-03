package initialize

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Tạo hàm kiểm tra lỗi. Nếu xảy ra lỗi thì nó sẽ thông báo chi tiết nơi xảy và và đồng thời lưu trong log
func CheckErrorPanic(err error, message string) {
	if err != nil {
		global.Logger.Error.Error("CheckErrorPanic: ", zap.String("message", message), zap.Error(err))
		panic(fmt.Sprintf("%s: %v", message, err)) // Panic để dừng chương trình nếu có lỗi nghiêm trọng, đồng thời cung cấp thông tin chi tiết về lỗi
	}
}

// Tạo hàm kiểm tra kết nối
func CheckConnection(db *sql.DB) {
	if err := db.Ping(); err != nil {
		CheckErrorPanic(err, "InitPostgreSQL: Failed to ping PostgreSQL")
	}
}

// Các hàm khởi tạo cho PostgreSQL sẽ được đặt ở đây
func InitPostgreSQL() {
	m := global.Config.PostgreSQL

	// Tạo chuỗi kết nối (DSN) cho PostgreSQL
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC"
	var stringConn = fmt.Sprintf(dsn, m.Host, m.User, m.Password, m.DBName, m.Port)

	db, err := sql.Open("postgres", stringConn)
	CheckErrorPanic(err, "InitPostgreSQL: Failed to connect to PostgreSQL")
	CheckConnection(db)

	global.Logger.Access.Info("InitPostgreSQL: Successfully connected to PostgreSQL")
	global.PostgreSQL = db // Gán *sql.DB vào global
	global.DB = database.New(global.PostgreSQL) // Khởi tạo database.Queries và gán vào global

	// Thiết lập connection pool
	SetPool()
	//genTableDAO()

	//Migrate tables
	//migrateTables()
}

// Hàm này sẽ thiết lập pool kết nối cho PostgreSQL
// Việc này thiết lập mở nhóm kết nối tối đa, số lượng kết nối nhàn rỗi tối đa và thời gian sống tối đa của kết nối
func SetPool() {
	p := global.Config.PostgreSQL

	global.PostgreSQL.SetMaxIdleConns(p.MaxIdleConns)                      // thiết lập số lượng kết nối nhàn rỗi tối đa để giữ sẵn sàng cho các yêu cầu mới mà không cần phải tạo kết nối mới từ đầu
	global.PostgreSQL.SetMaxOpenConns(p.MaxOpenConns)                      // thiết lập giới hạn số lượng kết nối tối đa để tránh quá tải cơ sở dữ liệu
	global.PostgreSQL.SetConnMaxIdleTime(time.Duration(p.MaxIdleConns))    // Thiết lập thời gian tối đa của một kết nối nhàn rỗi trước khi bị đóng
	global.PostgreSQL.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetime)) // SAu khi kết nối tồn tại hơn thời gian được thiết lập ở đây thì nó sẽ bị đóng vào loại khỏi pool
}