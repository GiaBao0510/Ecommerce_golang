package repository

import "context"

// Tạo Interface IUserRepo để định nghĩa các phương thức mà UserRepo sẽ triển khai
type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.)

}

