package authen

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/global"
	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

// Mô tả:
/*
	Tại chức năng cho refresh token, chúng ta sẽ thực hiện 1 số bước sau:
	- Bước 1 lấy access token và refresh token từ request
	- Bước 2 kiếm tra xem access token và refresh token có hợp lệ hay không, nếu không hợp lệ thì trả về lỗi
	- Bước 3 nếu access token hợp lệ thì xóa refresh token cũ và tạo mới access token và refresh token, trả về cho client
*/
type RefreshTokenUseCase struct {
	redis repository.IRedisRepository
	slog *loghelper.ServiceLogger
	userRepo  repository.IUserRepository
} 

func NewRefreshTokenUseCase(
	redis repository.IRedisRepository,
	slog *loghelper.ServiceLogger,
	userRepo repository.IUserRepository,
) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		redis: redis,
		slog: slog,
		userRepo: userRepo,
	}
}

func(r *RefreshTokenUseCase) RefreshToken(ctx context.Context, token *dto.Token) (*dto.Token, error){

	// ----- Các bước kiểm tra token -----

	// Kiểm tra xem refrestoken có tồn tại trong whitelist hay không, nếu có thì lấy uuid,nếu không tồn tại thì trả về lỗi
	refreshToken, err := r.redis.Get(ctx, _const.WhiteListRefreshToken + token.RefreshToken)
	if err != nil {
		return nil, apperrors.NewUnauthorizedError("Refresh token không hợp lệ hoặc đã hết hạn")
	}
	if refreshToken == "" {
		return nil, apperrors.NewUnauthorizedError("Refresh token không hợp lệ hoặc đã hết hạn")
	}

	// Chuyển đổi refreshToken từ []byte sang struct RefreshToken
	var refreshTokenData_old util.RefreshToken
	if err := json.Unmarshal([]byte(refreshToken), &refreshTokenData_old); err != nil {
		r.slog.LogError("Failed to unmarshal refresh token", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lấy thông tin user(user_id, role, email) dựa vào email
	userInfor, err := r.userRepo.UserVerificationInformationViaUID(ctx, refreshTokenData_old.UserID)
	if err != nil {
		r.slog.LogError("Failed to get user information", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	// Tạo access token mới và refresh token mới
	new_accessToken, err := util.GenerateAccessToken(userInfor.Uuid, userInfor.Email, int(userInfor.Role_id))
	if err != nil {
		r.slog.LogError("Failed to generate access token", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	new_refreshToken, err := util.GenerateRefreshToken(userInfor.Uuid)
	if err != nil {
		r.slog.LogError("Failed to generate refresh token", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	// Thu hồi Token bằng cách xóa refresh token cũ khỏi whitelist và thêm refresh token cũ vào blacklist
	if err := r.revokeRefreshToken(ctx, token.RefreshToken, userInfor.Uuid); err != nil {
		r.slog.LogError("Failed to revoke refresh token", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lưu access token mới và refresh token mới vào trong whitelist 
	if err := r.storeNewTokens(ctx, new_accessToken, new_refreshToken, userInfor.Uuid); err != nil {
		r.slog.LogError("Failed to store new tokens", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	return &dto.Token{
		AccessToken: new_accessToken,
		RefreshToken: new_refreshToken.Token,
	}, nil
}


// Thu hồi Token bằng cách xóa refresh token cũ khỏi whitelist và thêm refresh token cũ vào blacklist
func (r *RefreshTokenUseCase) revokeRefreshToken(ctx context.Context, refreshToken, userID string) error {

	whitelistKey := _const.WhiteListRefreshToken + refreshToken
    blacklistKey := _const.BlackList + refreshToken

	// Lấy thời gian hết hạn của refresh token từ cấu hình để đặt vào trong blacklist
	ttl, err := r.redis.GetTTL(ctx, whitelistKey)
	if err != nil {
		r.slog.LogError("Failed to get TTL of refresh token", err)
		return apperrors.NewInternalServerError(err)
	}

	// Đưa refresh token cũ bằng cách thêm vào blacklist
	if err := r.redis.Set(ctx, blacklistKey, userID, ttl); err != nil {
		r.slog.LogError("Failed to add old refresh token to blacklist", err)
		return apperrors.NewInternalServerError(err)
	}

	// Xóa refreshToken khỏi whitelist, nếu không tồn tại thì trả về lỗi
	if err := r.redis.Delete(ctx, whitelistKey); err != nil {
		r.slog.LogError("Failed to delete refresh token from whitelist", err)
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// Lưu token mới vào whitelist
func (r *RefreshTokenUseCase) storeNewTokens(
	ctx context.Context,
	newAccessToken string,
	newRefreshToken util.RefreshToken,
	userID string,
)  error {

	// Chuyển đổi refreshToken mới từ struct sang []byte
	refreshTokenData, err := json.Marshal(newRefreshToken)
	if err != nil {
		r.slog.LogError("Failed to marshal refresh token", err)
		return apperrors.NewInternalServerError(err)
	}

	// Lấy thông tin JTI từ access token mới, nếu không hợp lệ thì trả về lỗi
	jti, err := util.GetJTIFromClaims(newAccessToken)
	if err != nil {
		r.slog.LogError("Failed to get JTI from new access token", err)
		return apperrors.NewInternalServerError(err)
	}

	// Lưu refresh token (đã bị mã hóa) vào whitelist thông qua Redis với thời hạn là 7 ngày [Cấu trúc lưu trữ: Key: WhiteList_RefreshToken:<hashed_refresh_token>; Value: <user_id>]
	ttl := time.Duration(global.Config.Authentication.JWT.RefreshTokenExpirationDays) * 24 * time.Hour
	if err := r.redis.Set(ctx, _const.WhiteListRefreshToken + newRefreshToken.Token, refreshTokenData, ttl ); err != nil {
		r.slog.LogError("Failed to store refresh token in Redis", err, zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút
	ttl = time.Duration(global.Config.Authentication.JWT.AccessTokenExpirationMinutes) * time.Minute
	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút [Cấu trúc lưu trữ: Key: WhiteList_AccessToken:<jti>; Value: <user_id>]
	if err := r.redis.Set(ctx, _const.WhiteListAccessToken + jti, userID, ttl); err != nil {	
		r.slog.LogError("Failed to store access token in Redis", err, zap.Error(err))
		return apperrors.NewInternalServerError(err)
	}

	return nil
}
