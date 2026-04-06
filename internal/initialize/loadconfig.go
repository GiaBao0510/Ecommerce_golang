package initialize

import (
	"fmt"
	"log"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/spf13/viper"
)

// Tệp tin này chủ yếu để đọc các cấu hình trong file ở local
func LoadConfig(){
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
	my_viper.AddConfigPath("configs")
	my_viper.SetConfigName("local")
	my_viper.SetConfigType("yaml")
 
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