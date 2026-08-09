package repositoryimpl

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/mailjet/mailjet-apiv3-go"
	"go.uber.org/zap"
)


type EmailRepositoryImpl struct {
	mailjectClient *mailjet.Client
	logger         *loghelper.DBLogger
}

func NewEmailRepositoryImpl(
    mailjectClient *mailjet.Client,
    logger *loghelper.DBLogger,
) repository.IEmailRepository {
    return &EmailRepositoryImpl{
        mailjectClient: mailjectClient,
        logger:         logger,
    }
}

func (e *EmailRepositoryImpl) SendEmail(ctx context.Context, data models.EmailData) error {
	e.logger.LogInfo("Sending email via repository: ","",
		zap.String("to_email", data.ToEmail),
		zap.String("subject", data.Subject),
	)

	// Kiểm tra cấu hình Mailjet trước khi gửi email
	if err := e.checkMailjetConfig(); err != nil {
		return err
	}

	// Tạo email message
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: global.Config.Authentication.MailJet.From_mail,
				Name: global.Config.Authentication.MailJet.From_name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: data.ToEmail,
					Name: data.ToName,
				},
			},
			Subject: data.Subject,
			HTMLPart: data.HTMLBody,
			TextPart: data.TextBody,
		},
	}

	// Gửi email
	messages := mailjet.MessagesV31{Info: messagesInfo}

	// Gửi email và kiểm tra lỗi
	res, err := e.mailjectClient.SendMailV31(&messages)
	if err != nil {
		e.logger.LogError("Failed to send email via Mailjet: ",err,
			zap.Error(err),
			zap.String("to_email", data.ToEmail),
			zap.String("subject", data.Subject),
		)
		return err
	}

	e.logger.LogInfo("Email sent successfully via Mailjet: ","",
		zap.String("to_email", data.ToEmail),
		zap.String("subject", data.Subject),
		zap.Any("response", res),
	)

	return nil
}

func (e *EmailRepositoryImpl) checkMailjetConfig() error {

	if global.Config.Authentication.MailJet.API_key == "" || global.Config.Authentication.MailJet.Secret_key == "" || global.Config.Authentication.MailJet.From_mail == "" {
		e.logger.LogWarning(
			"Cấu hình Mailjet chưa đầy đủ. Vui lòng kiểm tra lại file cấu hình.", "",
			zap.Bool("api_key_set:", global.Config.Authentication.MailJet.API_key != ""),
			zap.Bool("secret_key_set:", global.Config.Authentication.MailJet.Secret_key != ""),
		)
		e.logger.LogInfo("Email có thể thiếu đi thông tin xác thực","",
			zap.String("api_key", global.Config.Authentication.MailJet.API_key),
			zap.String("secret_key", global.Config.Authentication.MailJet.Secret_key),
			zap.String("from_mail", global.Config.Authentication.MailJet.From_mail),
		)
	}
	return nil
}