package service

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
)

// Tạo interface IUserService để định nghĩa các phương thức mà UserService sẽ triển khai
type IUserService interface {
	Register(email string, purpose string) int
	//....
}

// UserService là struct triển khai interface IUserService
// Trong một service, chúng ta có thể chứa nhiều repository khác nhau để phục vụ cho các chức năng khác nhau của service đó
type UserService struct {
	userRepo repository.IUserRepository
	//....
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (obj *UserService) Register(email string, purpose string) int {

	// Nếu email đã tồn tại trong database, trả về lỗi
	if obj.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExisted
	}

	return response.ErrorCodeSuccess
}
