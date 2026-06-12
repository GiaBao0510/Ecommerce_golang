package logger

import (
	"os"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {

	loglevel := config.Loglevel
	// debug -> info -> warn -> error -> fatal -> panic

	var level zapcore.Level

	switch loglevel {
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

	// Lấy encoder tùy chỉnh
	encoder := getEncoderLog() 

	// Cấu hình hook để ghi log vào file với luân phiên (rotation)
	hook := lumberjack.Logger{
		Filename:   config.LogFile,    // Đường dẫn file log
		MaxSize:    config.MaxSize,    // Kích thước tối đa của file log (MB)
		MaxBackups: config.MaxBackups, // Số lượng file log cũ được giữ lại
		MaxAge:     config.MaxAge,     // Số ngày giữ lại file log cũ, Nếu quá thời gian này, file log sẽ bị xóa
		Compress:   config.Compress,   // Nén file log cũ thành định dạng .gz để tiết kiệm dung lượng
	}

	// Tạo zapcore với encoder, hook và level đã cấu hình
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	// Trả về LoggerZap với zap.Logger tùy chỉnh
	return &LoggerZap{zap.New(
		core,                               // Sử dụng core tùy chỉnh
		zap.AddCaller(),                    // Hiển thị tên file và số dòng gọi logger
		zap.AddStacktrace(zap.ErrorLevel)), // Chỉ ghi stacktrace cho log từ Error trở lên
	}
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
