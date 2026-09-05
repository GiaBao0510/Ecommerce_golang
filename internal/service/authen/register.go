package authen

import (
	//"context"

	//"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	userRepo repository.IUserRepository
	userRoleRepo repository.IUserRoleRepository
	redisRepo repository.IRedisRepository
	db *sql.DB
	slog *loghelper.ServiceLogger
	zapLogger *zap.Logger
}

func NewRegisterUseCase (
	db *sql.DB,
	logger *zap.Logger,
	userRepo repository.IUserRepository,
	userRoleRepo repository.IUserRoleRepository,
	redisRepo repository.IRedisRepository,
	
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
		redisRepo: redisRepo,
		userRoleRepo: userRoleRepo,
		db: db,
		zapLogger: logger,
		slog: loghelper.NewServiceLogger(logger, "RegisterUseCase"),
	}
}

func (r *RegisterUseCase) RegisterUser(ctx context.Context, input *models.CreateUsersRequest) error{

	// Check kiểm tra email có bị trùng lặp không
	checkDulicateEmail, err := r.userRepo.UserEmailExists(ctx, input.Email)
	if err != nil {
		r.slog.LogError("Quá trình kiểm tra email trùng lặp", err)
		return err
	}
	if checkDulicateEmail {
		r.slog.LogWarning(
			"Checking duplicate email failed",
			"Email already exists in the database",
			zap.String("email", input.Email),
		)
		return apperrors.NewEmailDuplicateError()
	}

	// check kiểm tra số điện thoại có bị trùng lặp không
	checkDulicatePhoneNum, err := r.userRepo.UserPhoneExists(ctx, input.Phone_num)
	if err != nil {
		r.slog.LogError("[Error] Quá trình kiểm tra số điện thoại trùng lặp", err)
		return err
	}
	if checkDulicatePhoneNum {
		r.slog.LogWarning(
			"Checking duplicate phone number failed",
			"Số điện thoại đã tồn tại trong cơ sở dữ liệu", 
			zap.String("phone_num", input.Phone_num),
		)
		return apperrors.NewPhoneDuplicateError()
	}

	// Băm mật khẩu trước khi lưu vào cơ sở dữ liệu
	hashPW, err := bcrypt.GenerateFromPassword([]byte(input.Password_hash), bcrypt.DefaultCost)
	if err != nil {
		r.slog.LogError("Failed to hash password", err, zap.Error(err))
		return err
	}
	input.Password_hash = string(hashPW)

	// Các thao tác trong transaction
	var newUUID string
 
	err = servicesupport.RunInTx(ctx, r.db, r.zapLogger, func(tx *sql.Tx) error {

		userRepoTx := r.userRepo.WithTx(tx)
		userRoleRepoTx := r.userRoleRepo.WithTx(tx)

		uid, err := userRepoTx.Create(ctx, input)
		if err != nil {
			return err //Rollback
		}
		newUUID = uid

		if _, err := userRoleRepoTx.Create(ctx, &models.UserRole{
			Id_role: 2,
			Uuid: uid,
		}); err != nil {
			return err //Rollback
		}

		return nil //Commit
	})

	if err != nil {
		r.slog.LogError("RegisterUser: Quá trình đăng ký người dùng thất bại", err, zap.String("email", input.Email ))
		return err
	}

	r.slog.LogInfo("User registration successful", "",
		zap.String("email", input.Email),
		zap.String("phone_num", input.Phone_num),
		zap.String("uuid", newUUID),
	)
	return nil
}