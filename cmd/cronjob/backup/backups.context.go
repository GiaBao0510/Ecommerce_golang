package backup

import "context"

// Context: sử dụng Backup strategy để thực hiện backup, có thể thay đổi strategy mà không cần thay đổi code trong cronjob
type BackupContext struct {
	strategy BackupStrategy
}

func NewBackupContext(strategy BackupStrategy) *BackupContext {
	return &BackupContext{
		strategy: strategy,
	}
}

//  SetStrategy cho phép đổi strategy tại runtime nếu sau này cần (ví dụ chọn theo config)
func (bc *BackupContext) SetStrategy(strategy BackupStrategy) {
	bc.strategy = strategy
}


// RunBackup sẽ gọi strategy hiện tại để thực hiện backup
func(bc *BackupContext) Run(
	ctx context.Context, 
	input string,
) (string, error) {
	return bc.strategy.Run(ctx, input)
}
