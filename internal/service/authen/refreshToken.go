package authen

import (
	"context"
	"encoding/json"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
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

func(r *RefreshTokenUseCase) RefreshToken(ctx context.Context, token string) (*dto.Token, error){

	// ----- Các bước kiểm tra token -----
	// Kiểm tra xem refrestoken có tồn tại trong whitelist hay không, nếu có thì lấy uuid,nếu không tồn tại thì trả về lỗi
	refreshToken, err := r.redis.Get(ctx, _const.WhiteListRefreshToken + token)
	if err != nil {
		return nil, apperrors.NewUnauthorizedError("Refresh token không hợp lệ hoặc đã hết hạn")
	}

	// Chuyển đổi refreshToken từ []byte sang struct RefreshToken
	var refreshTokenData util.RefreshToken

	if err := json.Unmarshal([]byte(refreshToken), &refreshTokenData); err != nil {
		r.slog.LogError("Failed to unmarshal refresh token", err)
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lấy thông tin user(user_id, role, email) dựa vào email
	userInfor, err := r.userRepo.UserVerificationInformationViaEmail(ctx, refreshTokenData.UserID)


	// Tạo access token mới và refresh token mới 
	
	// xóa refresh token cũ khỏi whitelist và thêm refresh token cũ vào blacklist

	// Lưu access token mới và refresh token mới vào trong whitelist 

	return nil, nil
}
