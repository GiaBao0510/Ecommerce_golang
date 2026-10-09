package email

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
)

type IEmailProviderService interface {
	// Gửi email qua nhà cung cấp email
	Send(ctx context.Context, message *dto.EmailMessage) error
	
	// Đóng kết nối với nhà cung cấp email
	Close() error
}