package backup

import (
	"context"
	"fmt"
	"os"
	"path/filepath"


	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

// Đây là phần Concrete Strategy, implement BackupStrategy interface để thực hiện backup với Cloudflare
type CloudflareBackupStrategy struct {}

//Hàm uploadToCloud_R2 sẽ upload file backup lên Cloudflare R2
func (c *CloudflareBackupStrategy) Run(ctx context.Context, filePath string) (string, error) {
	
	// Load cấu hình AWS từ biến môi trường
	cloudflareCpnfig := global.Config.Authentication.Cloudflare

	// Kiểm tra các biến môi trường cần thiết
	if cloudflareCpnfig.Endpoint == "" || 
		cloudflareCpnfig.SecretAccessKey == "" || 
		cloudflareCpnfig.BucketName == "" || 
		cloudflareCpnfig.AccessKeyID == "" {
		return "", fmt.Errorf("Missing Cloudflare R2 configuration in environment variables")
	}

	// Tạo cấu hình AWS với thông tin từ Cloudflare R2
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cloudflareCpnfig.AccessKeyID, 
				cloudflareCpnfig.SecretAccessKey, 
				"",
			),
		),
		config.WithRegion("auto"), // R2 dùng "auto" — không có khái niệm region như AWS thật
	)
	if err != nil {
		return "", apperrors.NewDetailedInternalServerError("Không load được cấu hình AWS SDK", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cloudflareCpnfig.Endpoint)
	})

	file, err := os.Open(filePath)
	if err != nil {
		return "", apperrors.NewDetailedInternalServerError("Không mở được file backup", err)
	}
	defer file.Close()

	uploader := manager.NewUploader(client)

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(cloudflareCpnfig.BucketName),
		Key:    aws.String(filepath.Base(filePath)),
		Body:   file,
	})
	if err != nil {
		global.Logger.Error.Error(
            "Không upload được file backup lên Cloudflare R2",
            zap.String("file", filePath),
            zap.Error(err),
        )

        return "", apperrors.NewDetailedInternalServerError(
            "Không upload được file backup lên Cloudflare R2",
            err,
        )
	}

	return filePath, nil
}