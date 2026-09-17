package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GiaBao0510/Ecommerce_golang/cmd/cronjob"
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/initialize"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {

	// Khởi động cấu hình
	initialize.LoadConfig()
	initialize.InitLogger()

	cronExpr := global.Config.CronJob.Backup_cron
	retentionDays := global.Config.CronJob.Backup_retention_days
	backupDir := global.Config.CronJob.Backup_dir

	// Kiểm tra các thông số cấu hình
	if cronExpr == "" {
		log.Fatal("Thiếu cấu hình cronExpr trong file config")
	}
	if retentionDays <= 0 {
		log.Fatal("Retention days phải lớn hơn 0")
	}
	if backupDir == "" {
		log.Fatal("Thiếu cấu hình backupDir trong file config")
	}

	// Tạo thư mục backup nếu chưa tồn tại
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		log.Fatalf("Không thể tạo thư mục backup %q: %v", backupDir, err)
	}

	// Lên lịch cron job để thực hiện backup
	c := cron.New()
	_, err := c.AddFunc(cronExpr, func() {
		cronjob.RunBackupJob(context.Background(), backupDir, retentionDays)
	})
	if err != nil {
		global.Logger.Error.Error("Lỗi khi thêm tác vụ cron: ", zap.Error(err))
		return
	}

	global.Logger.Access.Info("Cron job đã được lên lịch với biểu thức: ",
		zap.String("cronExpr", cronExpr),
		zap.Int("retentionDays", retentionDays),
		zap.String("backupDir", backupDir),
	)

	// Chạy thử trước 1 lần
	cronjob.RunBackupJob(context.Background(), backupDir, retentionDays)
	c.Start()

	// Lắng nghe SIGTERM/SIGINT để dừng cron job khi ứng dụng kết thúc
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.Logger.Access.Info("Nhận tín hiệu dừng, dừng cron job...")
	ctx := c.Stop() // trả về context để chờ các tác vụ đang chạy hoàn thành
	<-ctx.Done()    // chờ cho đến khi tất cả các tác vụ đang chạy hoàn thành
	global.Logger.Access.Info("Cron job đã dừng.")
}
