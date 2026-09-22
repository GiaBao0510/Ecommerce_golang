package oauth2

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"go.uber.org/zap"
)

func GenerateAccessTokenAndRefreshToken(
	slog *loghelper.ServiceLogger,
	redisRepo repository.IRedisRepository,
	ctx context.Context,
	userID string, 
	email string, 
	userRole int,
) (*models.LoginResponse, error) {

	// Tạo access token và refresh token
	accesstoken, err := util.GenerateAccessToken(userID, email, int(userRole))
	if err != nil {
		slog.LogError("Failed to generate access token", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}
	jti, err := util.GetJTIFromClaims(accesstoken)
	if err != nil {
		slog.LogError("Failed to get JTI from access token", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}
 
	// Tạo refresh token
	refreshToken, err := util.GenerateRefreshToken(userID)
	if err != nil {
		slog.LogError("Failed to generate refresh token", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}
	refreshTokenData, err := json.Marshal(refreshToken)
	if err != nil {
		slog.LogError("Failed to marshal refresh token", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lưu refresh token (đã bị mã hóa) vào whitelist thông qua Redis với thời hạn là 7 ngày [Cấu trúc lưu trữ: Key: WhiteList_RefreshToken:<hashed_refresh_token>; Value: <user_id>]
	ttl := time.Duration(global.Config.Authentication.JWT.RefreshTokenExpirationDays) * 24 * time.Hour
	if err := redisRepo.Set(ctx, _const.WhiteListRefreshToken + refreshToken.Token, refreshTokenData, ttl ); err != nil {
		slog.LogError("Failed to store refresh token in Redis", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút
	ttl = time.Duration(global.Config.Authentication.JWT.AccessTokenExpirationMinutes) * time.Minute
	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút [Cấu trúc lưu trữ: Key: WhiteList_AccessToken:<jti>; Value: <user_id>]
	if err := redisRepo.Set(ctx, _const.WhiteListAccessToken + jti, userID, ttl); err != nil {	
		slog.LogError("Failed to store access token in Redis", err, zap.Error(err))
		return nil, err
	}

	// Trả về access token và refresh token
	return &models.LoginResponse{
		AccessToken:  accesstoken,
		RefreshToken: refreshToken.Token,
	}, nil
}