package authen

import (
	//"context"

	//"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type RegisterUseCasee struct {
	userRepo repository.IUserRepository
	userRoleRepo repository.IUserRoleRepository
	redisRepo repository.IRedisRepository
	logger *loghelper.DBLogger
}

func NewRegisterUseCasee (
	userRepo repository.IUserRepository,
	userRoleRepo repository.IUserRoleRepository,
	redisRepo repository.IRedisRepository,
	logger *loghelper.DBLogger,
) *RegisterUseCasee {
	return &RegisterUseCasee{
		userRepo: userRepo,
		redisRepo: redisRepo,
		userRoleRepo: userRoleRepo,
		logger: logger,
	}
}

func (r *RegisterUseCasee) RegisterUser(ctx context.Context, input models.CreateUsersRequest) error{

	// Check kiểm tra email có bị trùng lặp không
	checkDulicateEmail, err := r.userRepo.UserEmailExists(ctx, input.Email)
	if err != nil {
		r.logger.LogError("Quá trình kiểm tra email trùng lặp", err)
		return err
	}
	if checkDulicateEmail {
		r.logger.LogWarning(
			"Checking duplicate email failed",
			"Email already exists in the database",
			zap.String("email", input.Email),
		)
		return nil
	}

	// check kiểm tra số điện thoại có bị trùng lặp không
	checkDulicatePhoneNum, err := r.userRepo.UserPhoneExists(ctx, input.Phone_num)
	if err != nil {
		r.logger.LogError("[Error] Quá trình kiểm tra số điện thoại trùng lặp", err)
		return err
	}
	if checkDulicatePhoneNum {
		r.logger.LogWarning(
			"Checking duplicate phone number failed",
			"Số điện thoại đã tồn tại trong cơ sở dữ liệu", 
			zap.String("phone_num", input.Phone_num),
		)
		return nil
	}

	// Tạo thông tin tài khoản người dùng mới
	result1, err := r.userRepo.Create(ctx, &input)
	if err != nil {
		r.logger.LogError("Quá trình tạo thông tin tài khoản người dùng mới", err)
		return err
	}

	if result1 == "" {
		r.logger.LogWarning(
			"create new user account failed",
			"Không thể tạo thông tin tài khoản người dùng mới", 
			zap.String("email", input.Email), 
			zap.String("phone_num", input.Phone_num),
		)
		return nil
	}

	// Tạo thông tin user với role
	result2, err := r.userRoleRepo.Create(ctx, &models.UserRole{Id_role: int32(2), Uuid: result1})
	if err != nil {
		r.logger.LogError("Quá trình tạo thông tin user với role", err)
		return err
	}
	if result2 != 0 {
		r.logger.LogWarning(
			"create user role failed",
			"Không thể tạo thông tin user với role", 
			zap.String("uuid", result1), 
			zap.Int32("role_id", 2),
		)
		return nil
	}

	r.logger.LogInfo("User registration successful", "",
		zap.String("email", input.Email),
		zap.String("phone_num", input.Phone_num),
		zap.String("uuid", result1),
	)
	return nil
}