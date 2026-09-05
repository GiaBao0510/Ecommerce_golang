package authen

import (
	"context"
	"time"

	_const "github.com/GiaBao0510/Ecommerce_golang/internal/const"
	"github.com/GiaBao0510/Ecommerce_golang/global"
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

func NewLoginUseCase(userRepo  repository.IUserRepository, redisRepo repository.IRedisRepository, slog *loghelper.ServiceLogger) *LoginUseCase {
	return &LoginUseCase{
		userRepo:  userRepo,
		redisRepo: redisRepo,
		slog:      slog,
	}
}

func (l *LoginUseCase) Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

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

func (l *LoginUseCase) loginByEmail(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {
	// Lấy thông tin người dùng bởi email
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaEmail(ctx, loginRequest.Account)
	if err != nil {
		l.slog.LogError("Login", err, zap.Error(err))
		return nil, err
	}

	return l.verifyUserCredentials(ctx, *userVeriInfor, loginRequest.Passoword)
}

func (l *LoginUseCase) loginByPhone(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {
	// lấy thông tin người dùng bởi phone
	userVeriInfor, err := l.userRepo.UserVerificationInformationViaPhone(ctx, loginRequest.Account)
	if err != nil {
		l.slog.LogError("Login", err, zap.Error(err))
		return nil, err
	}

	return l.verifyUserCredentials(ctx, *userVeriInfor, loginRequest.Passoword)
}

// Hàm kiểm tra thông tin xác thực của người dùng
func (l *LoginUseCase) verifyUserCredentials(ctx context.Context, userVeriInfor models.UserVerificationInformation, password string) (*models.LoginResponse, error) {
	// Kiểm tra xem mật khâủ đầu vào có khớp với mật khẩu đã băm không
	if err := bcrypt.CompareHashAndPassword([]byte(userVeriInfor.Password_hash), []byte(password)); err != nil {
		l.slog.LogWarning("Login", "Password is not match", zap.String("account", userVeriInfor.Email))
		return nil, apperrors.NewUnauthorizedError("Mật khẩu không hợp lệ")
	}

	// Kiểm tra xem trạng thái người dùng có hợp lệ không
	if userVeriInfor.Id_status == 2 || userVeriInfor.Id_status == 3 {
		l.slog.LogWarning("Login", "User is not active", zap.String("account", userVeriInfor.Email))
		return nil, apperrors.NewForbiddenError("Người dùng không hoạt động")
	}

	// Lấy thêm thông tin IP của người dùng từ 

	// Tạo access token và refresh token
	accesstoken, err := util.GenerateAccessToken(userVeriInfor.Uuid, userVeriInfor.Email, int(userVeriInfor.Role_id))
	if err != nil {
		l.slog.LogError("Failed to generate access token", err, zap.Error(err))
		return nil, err
	}
	jti, err := util.GetJTIFromClaims(accesstoken)
	if err != nil {
		l.slog.LogError("Failed to get JTI from access token", err, zap.Error(err))
		return nil, err
	}

	// Tạo refresh token
	refreshToken, err := util.GenerateRefreshToken()
	if err != nil {
		l.slog.LogError("Failed to generate refresh token", err, zap.Error(err))
		return nil, err
	}

	// Mã hóa refresh token trước khi lưu vào Redis
	hashedRefreshToken := util.HashToken(refreshToken)

	// Lưu refresh token (đã bị mã hóa) vào whitelist thông qua Redis với thời hạn là 7 ngày [Cấu trúc lưu trữ: Key: WhiteList_RefreshToken:<hashed_refresh_token>; Value: <user_id>]
	ttl := time.Duration(global.Config.Authentication.JWT.RefreshTokenExpirationDays) * 24 * time.Hour
	if err := l.redisRepo.Set(ctx, _const.WhiteListRefreshToken + ":" + hashedRefreshToken, userVeriInfor.Uuid, ttl); err != nil {
		l.slog.LogError("Failed to store refresh token in Redis", err, zap.Error(err))
		return nil, apperrors.NewInternalServerError(err)
	}

	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút
	ttl = time.Duration(global.Config.Authentication.JWT.AccessTokenExpirationMinutes) * time.Minute
	// Lưu access token vào whitelist thông qua Redis với thời hạn là 15 phút [Cấu trúc lưu trữ: Key: WhiteList_AccessToken:<jti>; Value: <user_id>]
	if err := l.redisRepo.Set(ctx, _const.WhiteListAccessToken+":"+jti, userVeriInfor.Uuid, ttl); err != nil {	
		l.slog.LogError("Failed to store access token in Redis", err, zap.Error(err))
		return nil, err
	}

	// Trả về access token và refresh token
	return &models.LoginResponse{
		AccessToken:  accesstoken,
		RefreshToken: refreshToken,
	}, nil
}
