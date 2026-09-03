package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"go.uber.org/zap"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type LogoutUseCase struct {
	redisRepo 	repository.IRedisRepository
	slog 		*loghelper.ServiceLogger
}

func NewLogoutUseCase(redisRepo repository.IRedisRepository, slog *loghelper.ServiceLogger) *LogoutUseCase{
	return &LogoutUseCase{
		redisRepo: redisRepo,
		slog: slog,
	}
}

func (l *LogoutUseCase) Logout(ctx context.Context, logoutRequest *models.LogoutRequest) error {
	
	// Lấy thời gian hết hạn của access token và refresh token từ Redis
	expirationTimeAccessToken, err := l.redisRepo.GetTTL(ctx, _const.WhiteListAccessToken+":"+logoutRequest.JTI)
	if err != nil {
		l.slog.LogError("Failed to get TTL for access token from Redis", err, zap.Error(err))
		return err
	}
	expirationTimeRefreshToken, err := l.redisRepo.GetTTL(ctx, _const.WhiteListRefreshToken+":"+logoutRequest.RefreshToken)
	if err != nil {
		l.slog.LogError("Failed to get TTL for refresh token from Redis", err, zap.Error(err))
		return err
	}

	// Kiểm tra xem thời gian hết hạn của token có hợp lệ không
	if expirationTimeAccessToken <= 0 || expirationTimeRefreshToken <= 0 {
		l.slog.LogWarning("Check TTL", "Token has already expired or does not exist in Redis")
		return apperrors.NewBadRequestError("Lỗi Token không hợp lệ") // Nếu token đã hết hạn hoặc không tồn
	} 

	// Kiểm tra xem token có tồn tại trong blacklist không
	accessTokenBeenBlackListed, err := l.redisRepo.Exists(ctx, _const.BlackList+":"+logoutRequest.JTI)
	if err != nil {
		l.slog.LogError("Failed to check if access token is blacklisted in Redis", err, zap.Error(err))
		return err
	}

	refreshTokenBeenBlackListed, err := l.redisRepo.Exists(ctx, _const.BlackList+":"+logoutRequest.RefreshToken)
	if err != nil {
		l.slog.LogError("Failed to check if refresh token is blacklisted in Redis", err, zap.Error(err))
		return err
	}

	// Nếu token đã tồn tại trong blacklist, trả về lỗi
	if accessTokenBeenBlackListed || refreshTokenBeenBlackListed {
		l.slog.LogWarning("Check Token in Blacklist", "Token has already been blacklisted")
		return apperrors.NewBadRequestError("Lỗi Token không hợp lệ") // Nếu token đã tồn tại trong blacklist
	}
	
	// Xóa token khỏi whitelist
	if err := l.redisRepo.Delete(ctx, _const.WhiteListAccessToken+":"+logoutRequest.JTI); err != nil {
		l.slog.LogError("Failed to delete access token from whitelist in Redis", err, zap.Error(err))
		return err
	} 
	if err := l.redisRepo.Delete(ctx, _const.WhiteListRefreshToken+":"+logoutRequest.RefreshToken); err != nil {
		l.slog.LogError("Failed to delete refresh token from whitelist in Redis", err, zap.Error(err))
		return err
	}

	// thêm token vào blacklist
	if err := l.redisRepo.Set(ctx, _const.BlackList+":"+logoutRequest.JTI, "1", expirationTimeAccessToken); err != nil {
		l.slog.LogError("Failed to add access token to blacklist in Redis", err, zap.Error(err))
		return err
	}
	if err := l.redisRepo.Set(ctx, _const.BlackList+":"+logoutRequest.RefreshToken, "1", expirationTimeRefreshToken); err != nil {
		l.slog.LogError("Failed to add refresh token to blacklist in Redis", err, zap.Error(err))
		return err
	}

	return nil
}