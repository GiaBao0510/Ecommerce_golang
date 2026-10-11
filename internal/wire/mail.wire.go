//go:build wireinject

package wire

import (
    "github.com/GiaBao0510/Ecommerce_golang/global"
    controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/mail"
    "github.com/GiaBao0510/Ecommerce_golang/internal/database"
    repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
    service "github.com/GiaBao0510/Ecommerce_golang/internal/service/email"
    "github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
    "github.com/google/wire"
    "go.uber.org/zap"
)

func InitMailRouterHandler(
    db *database.Queries,
    logger *zap.Logger,
) (controller.EmailControllerInterface, error) {  // ✅ Trả về interface
    wire.Build(
        // Helper providers
		NewDBLogger,
        NewMailServiceLogger,
        NewEmailProviderConfig,

        // Repository layer
        repositoryimpl.NewUserRepository,
        repositoryimpl.NewRedisRepositoryImpl,

        // Service layer
        service.NewProviderFactory,   // ✅ Factory tạo provider dựa trên config
        service.NewEmailService,      // ✅ Service dùng provider đã tạo

        // Controller layer
        controller.NewEmailController,
    )
    return nil, nil
}

// NewMailServiceLogger tạo ServiceLogger riêng cho module Email
// Khác với NewServiceLogger (dùng cho AuthFlow)
func NewMailServiceLogger(logger *zap.Logger) *loghelper.ServiceLogger {
    return loghelper.NewServiceLogger(logger, "EmailFlow")
}

// ĐÂY LÀ NƠI CHỌN PROVIDER:
// - Đọc global.Config.Authentication.EmailProvider → "mailtrap" hoặc "mailjet"
// - Dựa vào giá trị đó, lấy config tương ứng (API key, sender, ...)
// - Truyền vào NewProviderFactory() để tạo đúng provider
func NewEmailProviderConfig() service.ProviderConfig {

    // Đọc loại provider từ config YAML
    providerType := global.Config.Authentication.Email_provider

    switch service.ProviderType(providerType) {
    case service.ProviderMailtrap:
        // Lấy config của Mailtrap
        mtConfig := global.Config.Authentication.MailTrap
        return service.ProviderConfig{
            Type:        service.ProviderMailtrap,
            APIKey:      mtConfig.API_key,
            SandboxID:   mtConfig.SandboxID,
            SenderEmail: mtConfig.Sender_mail,
            SenderName:  mtConfig.Sender_name,
        }

    case service.ProviderMailjet:
        // Lấy config của Mailjet
        mjConfig := global.Config.Authentication.MailJet
        return service.ProviderConfig{
            Type:        service.ProviderMailjet,
            APIKey:      mjConfig.API_key,
            SecretKey:   mjConfig.Secret_key,
            SenderEmail: mjConfig.From_mail,
            SenderName:  mjConfig.From_name,
        }

    default:
        // Mặc định dùng Mailtrap nếu config sai
        mtConfig := global.Config.Authentication.MailTrap
        return service.ProviderConfig{
            Type:        service.ProviderMailtrap,
            APIKey:      mtConfig.API_key,
            SandboxID:   mtConfig.SandboxID,
            SenderEmail: mtConfig.Sender_mail,
            SenderName:  mtConfig.Sender_name,
        }
    }
}