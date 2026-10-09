package email

import (
	"context"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"

	"github.com/mailjet/mailjet-apiv3-go"
	"go.uber.org/zap"
)

type MailjetProvider struct {
	client *mailjet.Client
	senderMail string
	senderName string
	logger *zap.Logger
}

func NewMailjetProvider(
	config ProviderConfig, 
	logger *zap.Logger,
) (IEmailProviderService, error) {

	// Khởi tạo cấu hình
	client := mailjet.NewMailjetClient(
		config.APIKey,
		config.SecretKey,
	)

	return &MailjetProvider{
		client: client,
		senderMail: config.SenderEmail,
		senderName: config.SenderName,
		logger: logger,
	}, nil
}

// Thực hiện gửi
func(m *MailjetProvider) Send(ctx context.Context, message *dto.EmailMessage) error {

	// Tạo message
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.senderMail,
				Name: m.senderName,				
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: message.To,
				},
			},
			Subject: message.Subject,
			HTMLPart: message.Text,
		},
	}

	// thiết lập message
	messages := mailjet.MessagesV31{Info: messagesInfo}

	// Gửi email và kiểm tra lỗi
	_, err := m.client.SendMailV31(&messages)
	if err != nil {
		m.logger.Error("Lỗi khi gửi email qua Mailjet: ", zap.Error(err), zap.String("to", message.To))
		return err
	}

	return nil
}

func (m *MailjetProvider) Close() error {
	return nil
}