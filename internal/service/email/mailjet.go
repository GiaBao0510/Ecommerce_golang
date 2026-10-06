package email

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type MailjetConfig struct{
	APIKey string
	SecretKey string
	SenderMail string
	SenderName string
	AppURL string
}

type MailjetProvider struct {
	config *MailjetConfig
	logger *loghelper.ServiceLogger
}

func NewMailjetProvider(config *MailConfig) (IEmailProviderService, error) {
	return &MailjetProvider{
		config: &MailjetConfig{
			APIKey: global.Config.Authentication.,
		},
		logger: &loghelper.ServiceLogger{},
	}, nil
}

func(m *MailjetProvider) SendNotification(ctx context.Context, email *dto.Email) error{

	// Kiểm tra xem có tồn tại trong DB khônng
	return nil
}
func(m *MailjetProvider) SubmitAuthenticationInformation(ctx context.Context, email *dto.Email) error {

	// Kiểm tra xem có tồn tại không và đã xác thực chưa, nếu chưa thì gửi email xác thực
	return nil
}