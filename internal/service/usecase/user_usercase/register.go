package userusercase

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type RegisterUseCasee struct {
	userRepo repository.IUserRepository
	redisRepo repository.IRedisRepository
	logger *loghelper.DBLogger
}

func NewRegisterUseCasee (
	userRepo repository.IUserRepository,
	redisRepo repository.IRedisRepository,
	logger *loghelper.DBLogger,
) *RegisterUseCasee {
	return &RegisterUseCasee{
		userRepo: userRepo,
		redisRepo: redisRepo,
		logger: logger,
	}
}

func (r *RegisterUseCasee) RegisterUser(ctx context.Context, ) {

	// Check kiểm tra email có bị trùng lặp không
	// check kiểm tra số điện thoại có bị trùng lặp không

	// Nhận các thông tin
	// Thực hiện lưu thông tin người dùng vào redis với khóa key là 
}