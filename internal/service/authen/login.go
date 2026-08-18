package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type LoginUseCase struct {
	userRepo repository.IUserRepository
	redisRepo repository.IRedisRepository
	slog *loghelper.ServiceLogger
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (l *LoginUseCase) Login(ctx context.Context, loginRequest models.LoginRequest) (string, error) {
	
	// Kiểm tra đầu vào là email hay số điện thoại
	if util.DetectType(loginRequest.Account) == "email" {
		return l.loginByEmail(ctx, loginRequest)
	} else if util.DetectType(loginRequest.Account) == "phone" {
		return l.loginByPhone(ctx, loginRequest)
	}

	// Nếu không phải thì báo lỗi
	l.slog.LogWarning("Login", "Account is not valid", zap.String("account", loginRequest.Account))
	return "", apperrors.NewBadRequestError("Account không hợp lệ")
}

func (l *LoginUseCase) loginByEmail(ctx context.Context, loginRequest models.LoginRequest) (string, error) {
	// Lấy thông tin người dùng bởi email
	// Kiểm tra xem mật khâủ đầu vào có khớp với mật khẩu đã băm không
	// Kiểm tra xem trạng thái người dùng có hợp lệ không
	// Tạo access token và refresh token
	return "", nil
}

func (l *LoginUseCase) loginByPhone(ctx context.Context, loginRequest models.LoginRequest) (string, error) {
	// lấy thông tin người dùng bởi phone
	return "", nil
}

// Hàm kiểm tra thông tin xác thực của người dùng
func (l *LoginUseCase) verifyUserCredentials(ctx context.Context, account string, password string) (bool, error) {

}
