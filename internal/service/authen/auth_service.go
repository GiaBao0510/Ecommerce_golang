package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IAuthService interface {
	ChangePassword(ctx context.Context, uid string, newPassword string) error
	VerifyEmail(ctx context.Context, email, otp string) error
	VerifyPhone(ctx context.Context, phone, otp string) error
	Register(ctx context.Context, obj *models.CreateUsersRequest) error
	Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error)
	Logout(ctx context.Context, logoutReq *models.LogoutRequest) error
}

// Triển khai Interface IAuthService
type AuthService struct {
	registerUseCase   *RegisterUseCase
	loginUseCase      *LoginUseCase
	verifyUserUseCase *VerifyUserUsecase
	logoutUseCase     *LogoutUseCase
}

func NewAuthService(
	registerUseCase *RegisterUseCase,
	loginUseCase *LoginUseCase,
	verifyUserUseCase *VerifyUserUsecase,
	logoutUseCase *LogoutUseCase,
) IAuthService {
	return &AuthService{
		registerUseCase:   registerUseCase,
		loginUseCase:      loginUseCase,
		verifyUserUseCase: verifyUserUseCase,
		logoutUseCase:     logoutUseCase,
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
func (s *AuthService) Register(ctx context.Context, obj *models.CreateUsersRequest) error {
	return s.registerUseCase.RegisterUser(ctx, obj)
}
func (s *AuthService) Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {
	return s.loginUseCase.Login(ctx, loginRequest)
}
func (s *AuthService) Logout(ctx context.Context, logoutReq *models.LogoutRequest) error {
	return s.logoutUseCase.Logout(ctx, logoutReq)
}
