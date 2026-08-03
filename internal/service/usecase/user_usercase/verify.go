package userusercase

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type VerifyUserUsecase struct {
	userRepo  repository.IUserRepository
	redisRepo repository.IRedisRepository
	emailRepo repository.IEmailRepository
	logger    *loghelper.DBLogger
}

// Khởi động
func NewVerifyUserUsecase(
	userRepo repository.IUserRepository,
	emailRepo repository.IEmailRepository,
	redisRepo repository.IRedisRepository,
	logger *loghelper.DBLogger,
) VerifyUserUsecase {
	return VerifyUserUsecase{
		userRepo:  userRepo,
		redisRepo: redisRepo,
		logger:    logger,
		emailRepo: emailRepo,
	}
}

// Thực hiện xác thực email
func (u *VerifyUserUsecase) VerifyEmail(email string) error {

	// 1. kiểm tra xem email có tồn tại trong cơ sở dữ liệu hay không

	// 2. Nếu tồn tại, thực hiện các bước xác thực email
	// 2.1 Tạo mã OTP
	// 2.2 Lưu mã OTP vào Redis với thời gian hết hạn (5 phút)
	// 2.3 Tạo nội dung email với mã OTP
	// 2.4 Gửi email xác thực đến người dùng

}
