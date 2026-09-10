package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IAuthService interface {
	ChangePassword(ctx context.Context, uid string, newPassword string) error
	VerifyEmail(ctx context.Context, email, otp string) error
	VerifyPhone(ctx context.Context, phone, otp string) error
	Register(ctx context.Context, obj *models.CreateUsersRequest) error
	Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error)
	Logout(ctx context.Context, logoutReq *models.LogoutRequest) error
	RefreshToken(ctx context.Context, token *dto.Token) (*dto.Token, error)
}

// Triển khai Interface IAuthService
type AuthService struct {
	registerUseCase     *RegisterUseCase
	loginUseCase        *LoginUseCase
	verifyUserUseCase   *VerifyUserUsecase
	logoutUseCase       *LogoutUseCase
	refreshTokenUseCase *RefreshTokenUseCase
}

func NewAuthService(
	registerUseCase *RegisterUseCase,
	loginUseCase *LoginUseCase,
	verifyUserUseCase *VerifyUserUsecase,
	logoutUseCase *LogoutUseCase,
	refreshTokenUseCase *RefreshTokenUseCase,
) IAuthService {
	return &AuthService{
		registerUseCase:     registerUseCase,
		loginUseCase:        loginUseCase,
		verifyUserUseCase:   verifyUserUseCase,
		logoutUseCase:       logoutUseCase,
		refreshTokenUseCase: refreshTokenUseCase,
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
func (s *AuthService) RefreshToken(ctx context.Context, token *dto.Token) (*dto.Token, error) {
	return s.refreshTokenUseCase.RefreshToken(ctx, token)
}
