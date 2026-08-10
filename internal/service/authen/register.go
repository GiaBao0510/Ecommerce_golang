package authen

import (
	//"context"

	//"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	//"go.uber.org/zap"
)

type RegisterUseCasee struct {
	userRepo repository.IUserRepository
	//userRoleRepo
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

// func (r *RegisterUseCasee) RegisterUser(ctx context.Context, input models.RegisterRequest) error{

// 	// Check kiểm tra email có bị trùng lặp không
// 	checkDulicateEmail, err := r.userRepo.UserEmailExists(ctx, input.Email)
// 	if err != nil {
// 		r.logger.LogError("[Error] Quá trình kiểm tra email trùng lặp", err)
// 		return err
// 	}
// 	if checkDulicateEmail {
// 		r.logger.LogWarning("[Warning] Email đã tồn tại trong cơ sở dữ liệu", zap.String("email", input.Email))
// 		return nil
// 	}

// 	// check kiểm tra số điện thoại có bị trùng lặp không
// 	checkDulicatePhoneNum, err := r.userRepo.UserPhoneExists(ctx, input.Phone_num)
// 	if err != nil {
// 		r.logger.LogError("[Error] Quá trình kiểm tra số điện thoại trùng lặp", err)
// 		return err
// 	}
// 	if checkDulicatePhoneNum {
// 		r.logger.LogWarning("[Warning] Số điện thoại đã tồn tại trong cơ sở dữ liệu", zap.String("phone_num", input.Phone_num))
// 		return nil
// 	}


// 	// Nhận các thông tin
// 	// Thực hiện lưu thông tin người dùng vào redis với khóa key là 
// }