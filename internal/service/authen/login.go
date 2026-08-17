package authen

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
)

type LoginUseCase struct {
	userRepo repository.IUserRepository
	redisRepo repository.IRedisRepository
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (l *LoginUseCase) Login(ctx context.Context, email string, password string) (string, error) {
	// Kiểm tr đầu vào là email hay số điện thoại
	// Nếu là email thì gọi hàm login bằng email
	// Nếu là số điện thoại thì gọi hàm login bằng số điện thoại
	return "", nil
}

