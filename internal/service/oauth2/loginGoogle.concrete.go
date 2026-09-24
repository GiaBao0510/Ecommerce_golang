package oauth2

import (
	"context"
	"errors"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type LoginWithGoogleUseCase struct {
	userRepo     repository.IUserRepository
	userRoleRepo repository.IUserRoleRepository
	redisRepo    repository.IRedisRepository
	slog         *loghelper.ServiceLogger
}

func NewLoginWithGoogleUseCase(
	userRepo repository.IUserRepository,
	userRoleRepo repository.IUserRoleRepository,
	redisRepo repository.IRedisRepository,
	slog *loghelper.ServiceLogger,
) *LoginWithGoogleUseCase {
	return &LoginWithGoogleUseCase{
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		redisRepo:    redisRepo,
		slog:         slog,
	}
}

// triển khai đăng nhập thông qua Google OAuth2
func (l *LoginWithGoogleUseCase) Login(
	ctx context.Context,
	req *models.CreateUsersRequestNonStrict,
) (*models.LoginResponse, error) {

	// Kiểm tra xem email có trong DB không
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaEmail(ctx, req.Email)

	var uid string
	switch {

	// trong tường hợp đã tồn tạo
	case err == nil:
		uid = userVeriInfor.Uuid

	// Nếu không tìm thấy người dùng trong DB, thì tạo mới người dùng
	case errors.Is(err, apperrors.ErrNotFound):
		// Nếu không có trong DB thì tạo mới người dùng
		uid, err = l.userRepo.CreateUserFromOAuth2(ctx, req)
		if err != nil {
			l.slog.LogError("Lỗi tạo mới người dùng từ OAuth2", err, zap.Error(err))
			return nil, apperrors.NewDetailedInternalServerError("Lỗi tạo mới người dùng từ OAuth2", err)
		}

	// Lỗi thật
	default:
		l.slog.LogError("Lỗi truy xuất người dùng từ DB", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi truy xuất người dùng từ DB", err)
	}

	// truy xuất Role của người dùng dựa trên Uid
	roleID, err := l.userRoleRepo.GetRoleIDByUserID(ctx, uid)
	if err != nil {
		l.slog.LogError("Lỗi truy xuất Role của người dùng", err, zap.Error(err))
		return nil, apperrors.NewDetailedInternalServerError("Lỗi truy xuất Role của người dùng", err)
	}

	// Nếu có thì đăng nhập và trả về thông tin người dùng
	return GenerateAccessTokenAndRefreshToken(
		l.slog,
		l.redisRepo,
		ctx,
		uid,
		req.Email,
		int(roleID),
	)
}
