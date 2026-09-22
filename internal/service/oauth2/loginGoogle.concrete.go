package oauth2

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type LoginWithGoogleUseCase struct {
	userRepo repository.IUserRepository
	userRoleRepo repository.IUserRoleRepository
	redisRepo repository.IRedisRepository
	slog      *loghelper.ServiceLogger
}

func NewLoginWithGoogleUseCase(userRepo repository.IUserRepository, redisRepo repository.IRedisRepository, slog *loghelper.ServiceLogger) *LoginWithGoogleUseCase {
	return &LoginWithGoogleUseCase{
		userRepo: userRepo,
		redisRepo: redisRepo,
		slog:      slog,
	}
}

// triển khai đăng nhập thông qua Google OAuth2
func(l *LoginWithGoogleUseCase) Login(
	ctx context.Context, 
	req *models.CreateUsersRequestNonStrict,
) (*models.LoginResponse, error) {
	
	// Kiểm tra xem email có trong DB không
	
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaEmail(ctx, req.Email)
	if err != nil {
		l.slog.LogError("Lỗi kiểm tra thông tin người dùng", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi kiểm tra thông tin người dùng", err)
	}

	var Uid string

	// Nếu không thì tạo tài khoản người mới
	if userVeriInfor == nil {
		Uid, err = l.userRepo.CreateUserFromOAuth2(ctx, req)
		if err != nil {
			l.slog.LogError("Lỗi tạo tài khoản người dùng mới", err, zap.Error(err))
			return nil, apperrors.NewDetailedInternalServerError("Lỗi tạo tài khoản người dùng mới", err)
		}
	}

	// truy xuất Role của người dùng dựa trên Uid
	roleID, err := l.userRoleRepo.GetRoleIDByUserID(ctx, Uid)
	if err != nil {
		l.slog.LogError("Lỗi truy xuất Role của người dùng", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi truy xuất Role của người dùng", err)
	}

	
	// Nếu có thì đăng nhập và trả về thông tin người dùng
	
	
	return GenerateAccessTokenAndRefreshToken(
		l.slog,
		l.redisRepo,
		ctx,
		Uid, 
		req.Email, 
		int(roleID),
	)
}

