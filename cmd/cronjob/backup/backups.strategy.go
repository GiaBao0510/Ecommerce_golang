package backup

import "context"

// strategy là Interface chung cho tất cả các loại backup

type BackupStrategy interface {
	Run(ctx context.Context, path string) (string, error)
}