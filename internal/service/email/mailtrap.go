package email

import (
	"context"

	"github.com/mailtrap/mailtrap-go"
	"go.uber.org/zap"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
)

type MailtrapProvider struct {
	client *mailtrap.Client
	senderMail string
	senderName string
	logger *zap.Logger
}

func NewMailtrapProvider(
	config ProviderConfig,
	logger *zap.Logger,
) (IEmailProviderService, error) {

	// thực hiện việc khởi tạo cấu hình trước khi gửi
	client, err := mailtrap.NewClient(
		config.APIKey,
		mailtrap.WithSandbox(true),
		mailtrap.WithSandboxID(config.SandboxID), 
	)
	if err != nil {
		return nil, err
	}

	return &MailtrapProvider{
		client: client,
		senderMail: config.SenderEmail,
		senderName: config.SenderName,
		logger: logger,
	}, nil
}

func(m *MailtrapProvider) Send(ctx context.Context, message *dto.EmailMessage) error { 

	_, _, err := m.client.Send(
		ctx,
		&mailtrap.SendRequest{
			From: mailtrap.Address{
				Email: m.senderMail, 
				Name: m.senderName,
			},
			To: []mailtrap.Address{ 
				{Email: message.To,},
			},
			Subject: message.Subject,
			Category: message.Category,
			Text: message.Text,
		},
	)

	if err != nil {
		m.logger.Error(
            "Failed to send email via Mailtrap",
            zap.Error(err),
            zap.String("to", message.To),
        )
        return err
	}

	return nil
}

func (m *MailtrapProvider) Close() error {
	return nil
}