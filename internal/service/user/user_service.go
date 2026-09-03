package user

import (
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

// Tạo interface IUserService để định nghĩa các phương thức mà UserService sẽ triển khai
type IUserService interface {
	// CRUD
	GetByID(ctx context.Context, id string) (*models.Users, error)
	Create(ctx context.Context, obj *models.CreateUsersRequest) (string, error)
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
	slog *loghelper.ServiceLogger
}

// Constructor for UserService
func NewUserService(UserRepo repository.IUserRepository, logger *zap.Logger ) IUserService {
	return &UserService{
		UserRepo: UserRepo,
		slog: loghelper.NewServiceLogger(logger, "UserService"),
	}
}

//CRUD operations
func (s *UserService) GetByID(ctx context.Context, id string) (*models.Users, error) {

	return s.UserRepo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, obj *models.CreateUsersRequest) (string, error) {
		
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password_hash), bcrypt.DefaultCost)
	if err != nil {
		s.slog.LogError("Create", err, zap.Error(err))
		return "", apperrors.NewInternalServerError(err)
	}

	obj.Password_hash = string(hashedPassword)

	return s.UserRepo.Create(ctx, obj) 
}

func (s *UserService) Update_Put(ctx context.Context, id string, obj *models.UpdateUsersPutRequest) error {
	
	if servicesupport.RequireNonEmptyString(id, "User ID", "update_put", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không được để trống")
	}

	return s.UserRepo.Update_Put(ctx, id, obj)
}

func (s *UserService) Update_Patch(ctx context.Context, id string, obj *models.UpdateUsersPatchRequest) error {
	
	if servicesupport.RequireNonEmptyString(id, "User ID", "update_patch", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không được để trống")
	}

	return s.UserRepo.Update_Patch(ctx, id, obj)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	
	if servicesupport.RequireNonEmptyString(id, "User ID", "delete", s.slog) != nil {
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
	
	if servicesupport.RequireNonEmptyString(email, "Email", "get_user_by_email", s.slog) != nil {
		return nil, apperrors.NewBadRequestError("Email không được để trống")
	}
	
	return s.UserRepo.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error) {
	
	if servicesupport.RequireNonEmptyString(phone.String, "Phone", "get_user_by_phone", s.slog) != nil {
		return nil, apperrors.NewBadRequestError("Số điện thoại không được để trống")
	}
	
	return s.UserRepo.GetUserByPhone(ctx, phone)
}

// Update other operations
func (s *UserService) UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error {
	
	if servicesupport.RequireNonEmptyString(id, "User ID", "update_user_avatar", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không được để trống")
	}
	
	return s.UserRepo.UpdateUserAvatar_PATCH(ctx, id, avatarURL)
}