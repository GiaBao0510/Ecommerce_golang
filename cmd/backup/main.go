package main

import (
	"context"
	"os"

	"github.com/GiaBao0510/Ecommerce_golang/cmd/cronjob"
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {

	cronExpr := global.Config.CronJob.Backup_cron
	retentionDays := global.Config.CronJob.Backup_retention_days
	backupDir := global.Config.CronJob.Backup_dir

	if err := os.MkdirAll(backupDir, 0755); err != nil {
		global.Logger.Error.Error("Lỗi khi tạo thư mục backup: ", zap.Error(err))
	}

	c := cron.New()
	_, err := c.AddFunc(cronExpr, func(){
		cronjob.RunBackupJob(context.Background(),backupDir, retentionDays)
	})
	if err != nil {
		global.Logger.Error.Error("Lỗi khi thêm tác vụ cron: ", zap.Error(err))
		return
	}

	global.Logger.Access.Info("Cron job đã được lên lịch với biểu thức: ", zap.String("cronExpr", cronExpr))
	c.Start()
	select {} // Giữ cho chương trình chạy để cron job có thể thực hiện
}
