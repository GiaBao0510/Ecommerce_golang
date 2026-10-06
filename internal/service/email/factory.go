package email

import (
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
)

type ProviderType string

const (
	ProviderMailtrap ProviderType = "mailtrap"
	ProviderMailjet  ProviderType = "mailjet"
	ProviderMailgun  ProviderType = "mailgun"
	ProviderSendgrid ProviderType = "sendgrid"
)

type ProviderFactory interface {
	CreateProviver(config *MailConfig) (IEmailProviderService, error)
}

type MailtrapProviderFactory struct{}

func (f *MailtrapProviderFactory) CreateProviver(config *MailConfig) (IEmailProviderService, error) {
	return NewMailtrapProvider(config)
}

type MailjetProviderFactory struct{}

func (f *MailjetProviderFactory) CreateProviver(config *MailConfig) (IEmailProviderService, error) {
	return NewMailjetProvider(config)
}

func NewProviderFactory(providerType ProviderType) (ProviderFactory, error) {
	switch providerType {
	case ProviderMailtrap:
		return &MailtrapProviderFactory{}, nil
	case ProviderMailjet:
		return &MailjetProviderFactory{}, nil
	default:
		return nil, apperrors.NewDetailedInternalServerError(
			fmt.Sprintf("Unsupported email provider type: %s", providerType),
			apperrors.ErrInternalServer,
		)
	}
}