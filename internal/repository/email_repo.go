package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IEmailRepository interface {
	// Gửi email
	SendEmail(ctx context.Context, data models.EmailData) error 
}