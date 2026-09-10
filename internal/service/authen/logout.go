package authen

import (
	"context"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type LogoutUseCase struct {
	redisRepo 	repository.IRedisRepository
	slog 		  *loghelper.ServiceLogger
}

func NewLogoutUseCase(redisRepo repository.IRedisRepository, slog *loghelper.ServiceLogger) *LogoutUseCase{
	return &LogoutUseCase{
		redisRepo: redisRepo,
		slog: slog,
	}
}

func (l *LogoutUseCase) Logout(ctx context.Context, req *models.LogoutRequest) error {
	
	// Lấy JTI (JWT ID) từ access token
	jti, err := util.GetJTIFromClaims(req.AccessToken)
	if err != nil {
		l.slog.LogError("Failed to get JTI from access token", err, zap.Error(err))
		return apperrors.NewUnauthorizedError("Access token không hợp lệ")
	}

	// Lấy thời gian hết hạn của access token và refresh token từ Redis
	exp_AccessToken, err := l.redisRepo.GetTTL(ctx, _const.WhiteListAccessToken + jti)
	if err != nil {
		l.slog.LogError("Failed to get TTL for access token from Redis", err, zap.Error(err))
		return err
	}
	exp_RefreshToken, err := l.redisRepo.GetTTL(ctx, _const.WhiteListRefreshToken + req.RefreshToken)
	if err != nil {
		l.slog.LogError("Failed to get TTL for refresh token from Redis", err, zap.Error(err))
		return err
	}
	

	// Thực hiện thu hồi token bằng cách thêm vào blacklist trước, sau đó xóa khỏi whitelist

	err = l.revokeToken(
		ctx,
		_const.WhiteListAccessToken + jti,
		_const.WhiteListRefreshToken + req.RefreshToken,
		_const.BlackList + jti,
		_const.BlackList + req.RefreshToken,
		exp_AccessToken,
		exp_RefreshToken,
	)

	return nil
}

// Thực hiện quy trình thu hồi token. Quá trình thực hiện: Đặt vào blacklist trước, rồi xóa khỏi whitelist. Nếu có lỗi xảy ra trong quá trình thực hiện, sẽ trả về lỗi và không xóa khỏi whitelist.
func (l *LogoutUseCase) revokeToken(
	ctx context.Context,
	access_WhiteList string,
	refresh_WhiteList string,
	access_BlackList string,
	refresh_BlackList string,
	access_TTL time.Duration,
	refresh_TTL time.Duration,
) error {

	_, err := global.Redis.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		
		// 1. Thêm token vào blacklist trước
		pipe.Set(ctx,access_BlackList, "1", access_TTL)
		pipe.Set(ctx,refresh_BlackList, "1", refresh_TTL)

		// 2. Xóa token khỏi whitelist
		pipe.Del(ctx, access_WhiteList)
		pipe.Del(ctx, refresh_WhiteList)
		return nil
	})

	return err
}