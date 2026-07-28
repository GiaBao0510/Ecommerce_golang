package service

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)

// Tạo interface IUserService để định nghĩa các phương thức mà UserService sẽ triển khai
type IUserService interface {
	// CRUD
	GetByID(ctx context.Context, id string) (*models.Users, error)
	Create(ctx context.Context, obj *models.CreateUsersRequest) (int, error)
	Update_Put(ctx context.Context, id string, obj *models.UpdateUsersPutRequest) error
	Update_Patch(ctx context.Context, id string, obj *models.UpdateUsersPatchRequest) error
	Delete(ctx context.Context, id string) error

	// List operations 
	GetAll(ctx context.Context) ([]models.Users, error)

	// Search operations
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error)

	// Relationship operations

	// Update other operations
	UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error
}

type UserService struct {
	UserRepo repository.IUserRepository
	logger *zap.Logger
}

// Constructor for UserService
func NewUserService(UserRepo repository.IUserRepository, logger *zap.Logger ) IUserService {
	return &UserService{
		UserRepo: UserRepo,
		logger: logger,
	}
}

//CRUD operations
func (s *UserService) GetByID(ctx context.Context, id string) (*models.Users, error) {

	return s.UserRepo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, obj *models.CreateUsersRequest) (int, error) {

	return s.UserRepo.Create(ctx, obj)
}

func (s *UserService) Update_Put(ctx context.Context, id string, obj *models.UpdateUsersPutRequest) error {
	
	if id == "" {
		s.logger.Warn("Service: Update user - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "id is empty"),)
		return apperrors.NewBadRequestError("ID không được để trống")
	}

	return s.UserRepo.Update_Put(ctx, id, obj)
}

func (s *UserService) Update_Patch(ctx context.Context, id string, obj *models.UpdateUsersPatchRequest) error {
	
	if id == "" {
		s.logger.Warn("Service: Update user - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "id is empty"),)
		return apperrors.NewBadRequestError("ID không được để trống")
	}

	return s.UserRepo.Update_Patch(ctx, id, obj)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	
	if id == "" {
		s.logger.Warn("Service: Update user - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "id is empty"),)
		return apperrors.NewBadRequestError("ID không được để trống")
	}
	
	return s.UserRepo.Delete(ctx, id)
}

// List operations 
func (s *UserService) GetAll(ctx context.Context) ([]models.Users, error) {
	return s.UserRepo.GetAll(ctx)
}

// Search operations
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	
	if email == "" {
		s.logger.Warn("Service: GetUserByEmail - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "email is empty"),)
		return nil, apperrors.NewBadRequestError("Email không được để trống")
	}
	
	return s.UserRepo.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error) {
	
	if !phone.Valid {
		s.logger.Warn("Service: GetUserByPhone - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "phone is invalid"),)
		return nil, apperrors.NewBadRequestError("Số điện thoại không hợp lệ")
	}
	
	return s.UserRepo.GetUserByPhone(ctx, phone)
}

// Update other operations
func (s *UserService) UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error {
	
	if id == "" {
		s.logger.Warn("Service: UpdateUserAvatar_PATCH - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", "id is empty"),)
		return apperrors.NewBadRequestError("ID không được để trống")
	}
	
	return s.UserRepo.UpdateUserAvatar_PATCH(ctx, id, avatarURL)
}

// Xác minh đầu vào
func ValidateUserInput(serviceName, reason string, logger *zap.Logger, user *models.Users) error {

	if user.Email == "" {
		logger.Warn(serviceName + " - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", reason),)
		return apperrors.NewBadRequestError("Email không được để trống")
	}
	if user.Phone_num == "" {
		logger.Warn(serviceName + " - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", reason),)
		return apperrors.NewBadRequestError("Số điện thoại không được để trống")
	}
	if user.Address == "" {
		logger.Warn(serviceName + " - validation failed",
			zap.String("layer", "service"),
			zap.String("reason", reason),)
		return apperrors.NewBadRequestError("Địa chỉ không được để trống")
	}

	return nil
}