package email

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type MailTrapConfig struct{
	APIKey string
	SandboxID string
	SenderMail string
	SenderName string
}

type MailtrapProvider struct {
	config *MailTrapConfig
	logger *loghelper.ServiceLogger
}

func NewMailtrapProvider(config *MailConfig) (IEmailProviderService, error) {
	return &MailtrapProvider{
		config: &MailTrapConfig{
			APIKey: global.Config.Authentication.MailTrap.API_key,
			SandboxID: global.Config.Authentication.MailTrap.SandboxID,
			SenderMail: global.Config.Authentication.MailTrap.Sender_mail,
			SenderName: global.Config.Authentication.MailTrap.Sender_name,
		},
		logger: &loghelper.ServiceLogger{},
	}, nil
}

func(m *MailtrapProvider) SendNotification(ctx context.Context, email *dto.Email) error{

	// Kiểm tra xem có tồn tại trong DB khônng
	return nil
}
func(m *MailtrapProvider) SubmitAuthenticationInformation(ctx context.Context, email *dto.Email) error {

	// Kiểm tra xem có tồn tại không và đã xác thực chưa, nếu chưa thì gửi email xác thực
	return nil
}