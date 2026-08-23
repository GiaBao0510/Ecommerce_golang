package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/internal/util"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo  repository.IUserRepository
	redisRepo repository.IRedisRepository
	slog      *loghelper.ServiceLogger
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (l *LoginUseCase) Login(ctx context.Context, loginRequest models.LoginRequest) (*models.LoginResponse, error) {

	// Kiểm tra đầu vào là email hay số điện thoại
	if util.DetectType(loginRequest.Account) == "email" {
		return l.loginByEmail(ctx, loginRequest)
	} else if util.DetectType(loginRequest.Account) == "phone" {
		return l.loginByPhone(ctx, loginRequest)
	}

	// Nếu không phải thì báo lỗi
	l.slog.LogWarning("Login", "Account is not valid", zap.String("account", loginRequest.Account))
	return nil, apperrors.NewBadRequestError("Account không hợp lệ")
}

func (l *LoginUseCase) loginByEmail(ctx context.Context, loginRequest models.LoginRequest) (*models.LoginResponse, error) {
	// Lấy thông tin người dùng bởi email
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaEmail(ctx, loginRequest.Account)
	if err != nil {
		l.slog.LogError("Login", err, zap.Error(err))
		return nil, err
	}

	return l.verifyUserCredentials(ctx, userVeriInfor, loginRequest.Passoword)
}

func (l *LoginUseCase) loginByPhone(ctx context.Context, loginRequest models.LoginRequest) (*models.LoginResponse, error) {
	// lấy thông tin người dùng bởi phone
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaPhone(ctx, loginRequest.Account)
	if err != nil {
		l.slog.LogError("Login", err, zap.Error(err))
		return nil, err
	}

	return l.verifyUserCredentials(ctx, userVeriInfor, loginRequest.Passoword)
}

// Hàm kiểm tra thông tin xác thực của người dùng
func (l *LoginUseCase) verifyUserCredentials(ctx context.Context, userVeriInfor models.UserVerificationInformation, password string) (*models.LoginResponse, error) {
	// Kiểm tra xem mật khâủ đầu vào có khớp với mật khẩu đã băm không
	if err := bcrypt.CompareHashAndPassword([]byte(userVeriInfor.Password_hash), []byte(password)); err != nil {
		l.slog.LogWarning("Login", "Password is not match", zap.String("account", account))
		return nil, apperrors.NewUnauthorizedError("Mật khẩu không hợp lệ")
	}

	// Kiểm tra xem trạng thái người dùng có hợp lệ không
	if userVeriInfor.Id_status == 2 || userVeriInfor.Id_status == 3 {
		l.slog.LogWarning("Login", "User is not active", zap.String("account", userVeriInfor.Email))
		return nil, apperrors.NewForbiddenError("Người dùng không hoạt động")
	}

	// Tạo access token và refresh token
}
