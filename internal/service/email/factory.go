package email

import (
	"fmt"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)


// ProviderType định nghĩa các loại nhà cung cấp email
type ProviderType string

const (
	ProviderMailtrap ProviderType = "mailtrap"
	ProviderMailjet  ProviderType = "mailjet"
	ProviderMailgun  ProviderType = "mailgun"
	ProviderSendgrid ProviderType = "sendgrid"
)

// ProviderConfig chứa thông tin cấu hình cho nhà cung cấp email
type ProviderConfig struct {
	Type        ProviderType
	APIKey      string
	SecretKey   string
	SandboxID   int64
	SenderEmail string
	SenderName  string
	Timeout     time.Duration
}

// Factory trả về hàm khởi tạo đối với trường hợp cụ thể
func NewProviderFactory(
	config ProviderConfig,
	logger *zap.Logger,
) (IEmailProviderService, error) {
	switch config.Type {
	case ProviderMailtrap:
		return NewMailtrapProvider(config, logger)
	case ProviderMailjet:
		return NewMailjetProvider(config, logger)
	default:
		return nil, apperrors.NewDetailedInternalServerError(
			fmt.Sprintf("Unsupported email provider type: %s", config.Type),
			apperrors.ErrInternalServer,
		)
	}
}
