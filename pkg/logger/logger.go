package logger

import (
	"os"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// AppLoggers là wrapper bọc quanh zap.Logger
// Struct này để quản lý nhiều logger khác nhau trong cùng một ứng dụng, ví dụ:
// - Access logger: ghi log truy cập (request/response)
// - Error logger: ghi log lỗi (error, panic)
// - App logger: ghi log thông tin ứng dụng (info, debug)
type AppLoggers struct {
	Access  *zap.Logger
	Error   *zap.Logger
	App     *zap.Logger
	Warning *zap.Logger
}

// Khởi tạo logger dựa trên cấu hình
// - config.Loglevel: điều khiển mức log tối thiểu (debug/info/warn/error)
// - config.LogFormat: "json" hoặc "console" — xác định định dạng output
// - config.LogFile: đường dẫn file ghi log (sử dụng lumberjack để rotation)
func NewLogger(config setting.LoggerSetting) *AppLoggers {

	// -------------------------------------------------------
	// Bước 1: Chọn encoder (định dạng output) dựa trên LOG_FORMAT
	//
	// "json"    → dùng cho production, Loki, Elasticsearch đọc được
	// "console" → dùng cho development, có màu sắc, dễ đọc trực tiếp
	//
	// os.Getenv("LOG_FORMAT") đọc biến môi trường từ hệ thống
	// -------------------------------------------------------
	var encoder zapcore.Encoder
	logFormat := config.LogFormat
	if logFormat == "json" {
		encoder = getJSON_Encoder()
	} else {
		encoder = getConsole_Encoder()
	}

	// -------------------------------------------------------
	// Bước 2: Xác định log level từ config
	// Thứ tự ưu tiên: debug < info < warn < error < fatal
	// Ví dụ: nếu set "info", thì log DEBUG sẽ bị bỏ qua
	// -------------------------------------------------------
	var level zapcore.Level
	switch config.Loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// -------------------------------------------------------
	// Bước 3: Cấu hình lumberjack — thư viện quản lý file log
	// - File log sẽ ngày càng to lên nếu không có cơ chế xoay vòng (rotation)
	// - lumberjack tự động: cắt file khi đủ kích thước, xoá file cũ, nén file
	// -------------------------------------------------------

	// -------------------------------------------------------
	// Logger 1: Access Logger — ghi log truy cập (request/response)
	accessHook := &lumberjack.Logger{
		Filename:   config.LogAccessFile, // Đường dẫn file log
		MaxSize:    config.MaxSize,       // Kích thước tối đa của file log (MB)
		MaxBackups: config.MaxBackups,    // Số lượng file log cũ được giữ lại
		MaxAge:     config.MaxAge,        // Số ngày giữ lại file log cũ, Nếu quá thời gian này, file log sẽ bị xóa
		Compress:   config.Compress,      // Nén file log cũ thành định dạng .gz để tiết kiệm dung lượng
	}
	accessCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(accessHook),
		),
		zap.InfoLevel,
	)

	// -------------------------------------------------------
	// Logger 2: Error Logger — ghi log lỗi (errors)
	errorHook := &lumberjack.Logger{
		Filename:   config.LogErrorFile, // Đường dẫn file log
		MaxSize:    config.MaxSize,      // Kích thước tối đa của file log (MB)
		MaxBackups: config.MaxBackups,   // Số lượng file log cũ được giữ lại
		MaxAge:     config.MaxAge,       // Số ngày giữ lại file log cũ, Nếu quá thời gian này, file log sẽ bị xóa
		Compress:   config.Compress,     // Nén file log cũ thành định dạng .gz để tiết kiệm dung lượng
	}
	errorCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(errorHook),
		),
		zap.ErrorLevel,
	) 

	// -------------------------------------------------------
	// Logger 3: App Logger — ghi log app (request/response)
	appHook := &lumberjack.Logger{
		Filename:   config.LogAppFile, // Đường dẫn file log
		MaxSize:    config.MaxSize,    // Kích thước tối đa của file log (MB)
		MaxBackups: config.MaxBackups, // Số lượng file log cũ được giữ lại
		MaxAge:     config.MaxAge,     // Số ngày giữ lại file log cũ, Nếu quá thời gian này, file log sẽ bị xóa
		Compress:   config.Compress,   // Nén file log cũ thành định dạng .gz để tiết kiệm dung lượng
	}
	appCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(appHook),
		),
		level,
	) 

	// -------------------------------------------------------
	// Logger 4: Warning Logger — ghi log cảnh báo (warn)
	warningHook := &lumberjack.Logger{
		Filename: config.LogWarningFile,
		MaxSize: config.MaxSize,
		MaxAge: config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress: config.Compress,
	}
	warningCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(warningHook),
		),
		level,
	)

	// Trả về LoggerZap với zap.Logger tùy chỉnh
	return &AppLoggers{
		Access: zap.New(accessCore, zap.AddCaller()),
		Error: zap.New(errorCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		App: zap.New(appCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		Warning: zap.New(warningCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

// ============================================================
// getConsoleEncoder — Encoder cho môi trường DEVELOPMENT
// ============================================================
//
// Console encoder xuất ra dạng text có màu sắc, dễ đọc trên terminal.
// Ví dụ output:
//
//	2026-06-20T23:00:00.000+0700    INFO    main/main.go:10    Server started
//
// Định dạng này phù hợp khi:
//   - Chạy local bằng `go run` hoặc air
//   - Debug nhanh trực tiếp trên terminal
func getConsole_Encoder() zapcore.Encoder {

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
// getJSONEncoder — Encoder cho môi trường PRODUCTION
// ============================================================
//
// JSON encoder xuất ra từng dòng log là một JSON object hoàn chỉnh.
// Ví dụ output:
//
//	{"Time":"2026-06-20T23:00:00.000+0700","level":"info","caller":"main.go:10","msg":"Server started"}
//
// Định dạng này phù hợp để:
//   - Gửi vào Loki, Elasticsearch, Datadog
//   - Query bằng jq, LogQL, KQL
//   - Parse tự động bởi log aggregation tools
func getJSON_Encoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "Time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Trả về JSON Encoder — định dạng chuẩn cho production
	return zapcore.NewJSONEncoder(encoderConfig)
}
