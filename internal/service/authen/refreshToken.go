package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
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
}

func NewRefreshTokenUseCase(
	redis repository.IRedisRepository,
	slog *loghelper.ServiceLogger,
) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		redis: redis,
		slog: slog,
	}
}

func(r *RefreshTokenUseCase) RefreshToken(ctx context.Context, req *dto.Token) (*dto.Token, error){

	// Bước 1: Các bước xác minh access token
	// 1.1 Kiểm tra xem thuật toán JWT có hợp lệ hay không, nếu không hợp lệ thì trả về lỗi


	// 1.2 Kiểm tra xem chữ ký của access token có hợp lệ hay không, nếu không hợp lệ thì trả về lỗi
	// 1.3 Kiểm tra xem access token có bị hết hạn hay không, nếu hết hạn thì trả về lỗi
	// 1.4 Kiểm tra xem access token có bị blacklist hay không, nếu bị blacklist thì trả về lỗi

	// Bước 2: Các bước xác minh refresh token
	// 2.1 Kiểm tra xem refresh token có tồn tại trong whitelist hay không, nếu không tồn tại thì trả về lỗi
	// 2.2 Kiểm tra xem refresh token có bị hết hạn hay không, nếu hết hạn thì trả về lỗi
	// 2.3 Kiểm tra xem refresh token có bị blacklist hay không, nếu bị blacklist thì trả về lỗi

	// Bước 2: Xóa Access token cũ và refresh token cũ khỏi whitelist trong Redis
	// Bước 3: Thêm refresh token cũ vào blacklist trong Redis
	// Bước 4: Thêm access token mới và refresh token mới vào whitelist trong Redis

	return nil, nil
}
