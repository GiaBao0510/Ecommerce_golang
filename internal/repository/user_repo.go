package repository

// Tạo Interface IUserRepo để định nghĩa các phương thức mà UserRepo sẽ triển khai
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

// Triển khai Interface IUserRepo trong struct UserRepository
type UserRepository struct {

}

func (obj *UserRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}


