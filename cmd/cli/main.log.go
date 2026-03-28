package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ============================================================
// Hướng dẫn sử dụng thư viện Zap — Logging trong Golang
//
// Zap là thư viện logging hiệu suất cao của Uber dành cho Go.
// Nó nhanh hơn log chuẩn của Go và hỗ trợ structured logging,
// tức là log được gắn thêm các trường key-value có cấu trúc
// thay vì chỉ là một chuỗi văn bản thuần túy.
//
// Cài đặt:
//   go get go.uber.org/zap
//
// Các khái niệm Log Level (mức độ nhật ký):
//   Debug  — Thông tin chi tiết dùng khi phát triển/debug
//   Info   — Thông tin vận hành bình thường của ứng dụng
//   Warn   — Cảnh báo, chưa lỗi nhưng cần chú ý
//   Error  — Lỗi xảy ra nhưng ứng dụng vẫn tiếp tục chạy
//   Fatal  — Lỗi nghiêm trọng, gọi os.Exit(1) sau khi log
// ============================================================

func main() {

	// =========================================================
	// PHẦN 1: Sugar Logger — Cách log đơn giản, linh hoạt
	// =========================================================
	//
	// Zap cung cấp hai kiểu logger:
	//
	//   Logger      — Nhanh hơn, yêu cầu khai báo kiểu dữ liệu
	//                 rõ ràng cho từng trường (zap.String, zap.Int, ...)
	//
	//   SugaredLogger — Chậm hơn một chút, nhưng linh hoạt hơn,
	//                   dùng cú pháp tương tự fmt.Printf / fmt.Println.
	//                   Phù hợp khi không cần tối ưu hiệu năng tối đa.
	//
	// Gọi .Sugar() trên một Logger để lấy SugaredLogger.

	sugarLogger := zap.NewExample().Sugar()

	// Infof: định dạng chuỗi như fmt.Printf
	sugarLogger.Infof("Ứng dụng đang chạy, phiên bản: %s", "1.0.0")

	// Infow: log kèm các cặp key-value (loosely typed)
	sugarLogger.Infow("Kết nối thành công",
		"host", "localhost",
		"port", 8080,
	)

	// =========================================================
	// PHẦN 2: Logger thường — Nhanh hơn, structured rõ ràng
	// =========================================================
	//
	// Với Logger (không phải Sugar), mỗi trường phải được khai
	// báo kiểu dữ liệu cụ thể bằng các hàm zap.String(), zap.Int(),...
	// Điều này giúp Zap tối ưu hóa hiệu năng tốt hơn.

	logger := zap.NewExample()
	logger.Info("Khởi động server",
		zap.String("version", "1.0.0"),
		zap.Int("port", 8080),
	)

	// =========================================================
	// PHẦN 3: Ba preset (cấu hình sẵn) của Zap
	// =========================================================
	//
	// Zap cung cấp 3 cấu hình có sẵn để dùng nhanh mà không
	// cần tự tay cấu hình. Mỗi preset có log level và định dạng
	// đầu ra khác nhau:
	//
	//   zap.NewExample()
	//     - Log level: Debug trở lên
	//     - Định dạng: JSON, không có timestamp
	//     - Mục đích: Dùng trong ví dụ, demo, playground
	//
	//   zap.NewDevelopment()
	//     - Log level: Debug trở lên
	//     - Định dạng: Console thân thiện (có màu, có caller)
	//     - Mục đích: Dùng khi đang phát triển ứng dụng
	//
	//   zap.NewProduction()
	//     - Log level: Info trở lên (bỏ qua Debug)
	//     - Định dạng: JSON chuẩn (dễ parse bằng các công cụ log)
	//     - Mục đích: Dùng trên môi trường production thực tế

	// --- Example ---
	exampleLogger := zap.NewExample()
	exampleLogger.Debug("Đây là log từ NewExample()")

	// --- Development ---
	devLogger, _ := zap.NewDevelopment()
	devLogger.Debug("Đây là log từ NewDevelopment() — thấy cả file:line")

	// --- Production ---
	prodLogger, _ := zap.NewProduction()
	prodLogger.Info("Đây là log từ NewProduction() — định dạng JSON")
	// prodLogger.Debug(...)  <-- sẽ KHÔNG hiển thị vì level mặc định là Info

	// =========================================================
	// PHẦN 4: Tự tùy chỉnh Logger với zapcore
	// =========================================================
	//
	// Khi 3 preset trên không đủ linh hoạt (ví dụ: bạn muốn ghi
	// log ra cả console lẫn file, hoặc muốn định dạng riêng),
	// bạn có thể tự xây dựng logger từ các thành phần của zapcore.
	//
	// Một logger tùy chỉnh gồm 3 thành phần chính:
	//
	//   Encoder  — Quyết định định dạng log (JSON hay Console,
	//              timestamp, màu sắc, ...)
	//
	//   WriteSyncer — Quyết định log được ghi đi đâu
	//                 (console, file, hoặc cả hai)
	//
	//   Level    — Chỉ những log từ level này trở lên mới được ghi
	//
	// zapcore.NewCore(encoder, writeSyncer, level) kết hợp 3 thứ
	// trên thành một "core" — đây là trái tim của logger tùy chỉnh.

	encoder := getEncoderLog()     // Bước 1: Định nghĩa định dạng log
	writeSyncer := getWriterSync() // Bước 2: Định nghĩa nơi ghi log

	// Bước 3: Kết hợp thành core, chỉ ghi log từ InfoLevel trở lên
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	// Bước 4: Tạo logger từ core
	// zap.AddCaller() giúp log hiển thị thêm tên file và số dòng gọi logger
	customLogger := zap.New(core, zap.AddCaller())

	customLogger.Info("Log dòng 1 — ghi ra cả console lẫn file")
	customLogger.Info("Log dòng 2 — ghi ra cả console lẫn file")
	customLogger.Warn("Đây là cảnh báo — cũng được ghi vì Warn >= Info")
	// customLogger.Debug(...) <-- sẽ KHÔNG ghi vì Debug < Info
}

// ============================================================
// getEncoderLog — Định nghĩa định dạng (format) của log
// ============================================================
//
// Encoder kiểm soát mỗi dòng log trông như thế nào:
// timestamp ở đâu, level hiển thị kiểu gì, caller có hiện không...
//
// Ở đây ta dùng NewConsoleEncoder (định dạng dễ cho người đọc),
// thay vì NewJSONEncoder (định dạng máy đọc, dùng trên production).
func getEncoderLog() zapcore.Encoder {

	// Bắt đầu từ cấu hình chuẩn của Production làm nền tảng
	encoderConfig := zap.NewProductionEncoderConfig()

	// Thay đổi định dạng timestamp sang ISO 8601
	// Ví dụ: 2024-01-15T10:30:00.000+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Đổi tên key của trường thời gian từ "ts" (mặc định) sang "Time"
	// để dễ đọc hơn trong output
	encoderConfig.TimeKey = "Time"

	// Hiển thị log level bằng chữ HOA có màu (INFO, WARN, ERROR,...)
	// CapitalColorLevelEncoder chỉ có màu khi output là terminal
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Hiển thị caller theo dạng ngắn gọn, ví dụ: main/main_log.go:120
	// thay vì đường dẫn đầy đủ tuyệt đối
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Trả về Console Encoder — định dạng dễ cho người đọc
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// ============================================================
// getWriterSync — Định nghĩa nơi ghi log (output destination)
// ============================================================
//
// WriteSyncer là đích đến của log. Ở đây ta ghi log ra hai nơi
// cùng lúc: terminal (stderr) và file — nhờ NewMultiWriteSyncer.
func getWriterSync() zapcore.WriteSyncer {

	// Tạo/mở file log để ghi
	// os.OpenFile nhận 3 tham số:
	//   1. Đường dẫn file — sẽ tạo mới nếu chưa tồn tại (nhờ O_CREATE)
	//   2. Flag — cách mở file:
	//        os.O_CREATE  : Tạo file nếu chưa có
	//        os.O_WRONLY  : Chỉ ghi (không đọc)
	//        os.O_APPEND  : Ghi nối vào cuối file (không ghi đè)
	//   3. Permission (quyền truy cập) — 0666 nghĩa là:
	//        owner, group, other đều có quyền đọc và ghi
	//        (rw-rw-rw-), thường được hệ thống umask thu hẹp lại)
	//
	// Lưu ý: Bạn cần tạo thư mục "logs/" trước khi chạy,
	// hoặc dùng os.MkdirAll("logs", 0755) để tự tạo.
	file, err := os.OpenFile(
		"logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal("Không thể mở file log: ", err)
	}

	// zapcore.AddSync bọc file/writer thành WriteSyncer —
	// đảm bảo log được flush (xả bộ đệm) an toàn, tránh mất dữ liệu.
	syncFile := zapcore.AddSync(file)

	// os.Stderr là luồng lỗi tiêu chuẩn — thường dùng để ghi log
	// ra terminal thay vì os.Stdout, để tách biệt với output chương trình.
	syncConsole := zapcore.AddSync(os.Stderr)

	// NewMultiWriteSyncer cho phép ghi log ra nhiều đích cùng lúc.
	// Ở đây: vừa ra console, vừa lưu vào file.
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
