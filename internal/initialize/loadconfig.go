package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Tạo một hàm load cấu hình từ .env
func loadEnvConfig() {
	// Load file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}
}

// Tệp tin này chủ yếu để đọc các cấu hình trong file ở local
func LoadConfig(){

	// Load biến môi trường từ file .env
	loadEnvConfig()

	// Lấy các thông tin từ biến môi trường 
	configPath := os.Getenv("ConfigPath")
	configName := os.Getenv("ConfigName")
	configType := os.Getenv("ConfigType")

	// Kiểm tra các thông tin cấu hình từ biến môi trường
	if configPath == "" || configName == "" || configType == "" {
		log.Fatal("ERROR: Thiếu thông tin cấu hình trong biến môi trường: ConfigPath, ConfigName, ConfigType")
	}

	// ----------------------------------------------------------
	// BƯỚC 1: Khởi tạo một instance Viper riêng (thay vì dùng
	// global). Cách này giúp dễ quản lý khi project có nhiều
	// file config khác nhau (vd: local, production, test...).
	// ----------------------------------------------------------
	my_viper := viper.New()
 
	// ----------------------------------------------------------
	// BƯỚC 2: Cấu hình đường dẫn & định dạng file config.
	//   - AddConfigPath : thư mục chứa file config
	//   - SetConfigName : tên file (KHÔNG bao gồm phần mở rộng)
	//   - SetConfigType : định dạng file (yaml, json, toml, ...)
	//
	// Với cấu hình dưới đây, Viper sẽ tìm file: configs/local.yaml
	// ----------------------------------------------------------
	my_viper.AddConfigPath(configPath)
	my_viper.SetConfigName(configName)
	my_viper.SetConfigType(configType)
 
	// ----------------------------------------------------------
	// BƯỚC 3: Đặt giá trị mặc định (SetDefault).
	// Nếu key không tồn tại trong file YAML, Viper sẽ dùng
	// giá trị mặc định này thay vì trả về zero-value.
	// ----------------------------------------------------------
	my_viper.SetDefault("server.port", 8080)
	my_viper.SetDefault("security.jwt.key", "default-secret-key")
 
	// ----------------------------------------------------------
	// BƯỚC 4: Đọc nội dung file config vào bộ nhớ.
	// Nếu file không tìm thấy hoặc bị lỗi cú pháp YAML,
	// chương trình sẽ dừng ngay tại đây với thông báo lỗi.
	// ----------------------------------------------------------
	if err := my_viper.ReadInConfig(); err != nil {
		log.Fatal("Lỗi khi đọc file config: ", err)
	}

	// Map cấu hình vào struct global.Config đã định nghĩa sẵn. Nếu có lỗi trong quá trình unmarshal
	if err := my_viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Lỗi khi unmarshal config: %v\n", err)
	}
 
	fmt.Println("✅ Đọc file config thành công:", my_viper.ConfigFileUsed())
	fmt.Println()	
}