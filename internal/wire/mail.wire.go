// internal/wire/status.wire.go
//go:build wireinject

package wire

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/mail"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/email"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitMailRouterHandler(
	db *database.Queries, 
	logger *zap.Logger,
) (*controller.EmailController, error) {
	wire.Build(
		// Helper Provider
		NewMailServiceLogger,

		// Repository
		repositoryimpl.NewUserRepository,
		repositoryimpl.NewRedisRepositoryImpl,

		// Service
		service.NewProviderFactory,
		service.NewEmailService,
		
		// Controller
		controller.NewEmailController,
		
	)
	return nil, nil
}

// NewMailServiceLogger tạo một ServiceLogger cho dịch vụ email
func NewMailServiceLogger(logger *zap.Logger) *loghelper.ServiceLogger {
	return loghelper.NewServiceLogger(logger, "EmailService")
}

// NewEmailProviderConfig đọc cấu hình từ global config
// và tạo ProviderConfig phù hợp
//
// ĐÂY LÀ NƠI CHỌN PROVIDER:
// - Đọc global.Config.Authentication.EmailProvider → "mailtrap" hoặc "mailjet"
// - Dựa vào giá trị đó, lấy config tương ứng (API key, sender, ...)
// - Truyền vào NewProviderFactory() để tạo đúng provider
func NewEmailProviderConfig