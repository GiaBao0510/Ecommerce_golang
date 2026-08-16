package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IAuthService interface {
	ChangePassword(ctx context.Context, uid string, newPassword string) error
	VerifyEmail(ctx context.Context, email, otp string) error
	VerifyPhone(ctx context.Context, phone, otp string) error
	Register(ctx context.Context,obj models.CreateUsersRequest) error
	Login(ctx context.Context, email string, password string) (string, error)
}

// Triển khai Interface IAuthService
type AuthService struct {
	registerUseCase RegisterUseCase
	loginUseCase LoginUseCase
	verifyUserUseCase VerifyUserUsecase
}

func NewAuthService(
	registerUseCase RegisterUseCase,
	loginUseCase LoginUseCase,
	verifyUserUseCase VerifyUserUsecase,
) IAuthService {
	return &AuthService{
		registerUseCase: registerUseCase,
		loginUseCase: loginUseCase,
		verifyUserUseCase: verifyUserUseCase,
	}
}

func (s *AuthService) ChangePassword(ctx context.Context, uid string, newPassword string) error {
	return s.verifyUserUseCase.ChangePassword(ctx, uid, newPassword)
}
func (s *AuthService) VerifyEmail(ctx context.Context, email, otp string) error {
	return s.verifyUserUseCase.VerifyEmail(ctx, email, otp)
}
func (s *AuthService) VerifyPhone(ctx context.Context, phone string, otp string) error {
	return s.verifyUserUseCase.VerifyPhone(ctx, phone, otp)
}
func (s *AuthService) Register(ctx context.Context, obj models.CreateUsersRequest) error {
	return s.registerUseCase.RegisterUser(ctx, obj)
}
func (s *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	return s.loginUseCase.Login(ctx, email, password)
}