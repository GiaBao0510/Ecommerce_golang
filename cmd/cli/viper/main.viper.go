package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ============================================================
// Config là struct dùng để ánh xạ (map) dữ liệu từ file YAML
// vào các biến Go có kiểu dữ liệu rõ ràng.
//
// Tag `mapstructure:"..."` phải khớp với key trong file YAML.
// ============================================================
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {

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

	fmt.Println("✅ Đọc file config thành công:", my_viper.ConfigFileUsed())
	fmt.Println()

	// ----------------------------------------------------------
	// BƯỚC 5: Bật tính năng theo dõi thay đổi file (hot-reload).
	// Khi file YAML được chỉnh sửa và lưu, callback bên dưới
	// sẽ tự động chạy — KHÔNG cần khởi động lại chương trình.
	// Rất hữu ích trong môi trường phát triển (development).
	// ----------------------------------------------------------
	my_viper.WatchConfig()
	my_viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("⚡ File config vừa thay đổi:", e.Name)
		// Thực tế: reload lại biến config ở đây
	})

	// ----------------------------------------------------------
	// BƯỚC 6: Đọc từng giá trị riêng lẻ bằng getter của Viper.
	// Viper hỗ trợ nhiều kiểu: GetInt, GetString, GetBool,
	// GetFloat64, GetDuration, GetStringSlice, ...
	// Key được viết theo dạng "cha.con" (dot notation).
	// ----------------------------------------------------------
	fmt.Println("=== Đọc từng giá trị (Get) ===")
	fmt.Println("Server port  :", my_viper.GetInt("server.port"))
	fmt.Println("JWT key      :", my_viper.GetString("security.jwt.key"))
	fmt.Println()

	// ----------------------------------------------------------
	// BƯỚC 7: Unmarshal — ánh xạ toàn bộ config vào struct.
	// Đây là cách được khuyến nghị khi cần dùng nhiều field,
	// giúp code gọn hơn so với việc gọi Get() từng dòng.
	// ----------------------------------------------------------
	var config Config

	if err := my_viper.Unmarshal(&config); err != nil {
		// Lưu ý: phải dùng fmt.Printf (không phải fmt.Println)
		// khi chuỗi có chứa ký tự định dạng như %v, %s, %d...
		log.Fatalf("Không thể ánh xạ config vào struct: %v", err)
	}

	fmt.Println("=== Đọc qua struct (Unmarshal) ===")
	fmt.Println("Config port  :", config.Server.Port)
	fmt.Println("Config JWT   :", config.Security.JWT.Key)
	fmt.Println()

	// ----------------------------------------------------------
	// BƯỚC 8: Duyệt slice — đọc danh sách nhiều database.
	// Trong YAML, đây là một mảng các object (dùng dấu "-").
	// Sau khi Unmarshal, ta có thể duyệt bằng vòng lặp for-range.
	// ----------------------------------------------------------
	fmt.Println("=== Danh sách Database ===")
	fmt.Println("Số lượng database:", len(config.Databases))

	for i, db := range config.Databases {
		fmt.Printf("  [%d] user=%-10s password=%-10s host=%s\n",
			i+1, db.User, db.Password, db.Host)
	}
}
