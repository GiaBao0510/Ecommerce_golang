package email

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
)

type IEmailProviderService interface {
	SendNotification(ctx context.Context, email *dto.Email) error
	SubmitAuthenticationInformation(ctx context.Context, email *dto.Email) error
}