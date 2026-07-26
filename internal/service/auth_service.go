package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IAuthService interface {
	ChangePassword(ctx context.Context, uid string, newPassword string) error
	VerifyEmail(ctx context.Context, uid string) error
	VerifyPhone(ctx context.Context, uid string) error
	Register(ctx context.Context,obj *models.Users) (int, error)
	Login(ctx context.Context, email string, password string) (string, error)
}
