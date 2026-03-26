package service

import "github.com/GiaBao0510/Ecommerce_golang/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

// Hàm khởi tạo mới cho UserService
func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

func (obj *UserService) GetInfoService() string {
	return obj.UserRepo.GetInfo()
}
