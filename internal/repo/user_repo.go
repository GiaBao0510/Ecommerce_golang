package repo

type UserRepo struct {
	
}

// Hàm khởi tạo mới cho UserRepo
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (obj *UserRepo) GetInfo() string {
	return "UserRepo: GetInfo"
}

