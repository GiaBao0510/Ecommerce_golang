package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"go.uber.org/zap"
)

type IUserRoleService interface {
	Create(ctx context.Context, obj *models.UserRole) (int, error)
	Update(ctx context.Context, id string, obj *models.UserRole) error
	Delete(ctx context.Context, uuid string, roleID int32) error
	GetUserByRoleID(ctx context.Context, roleID int32) ([]models.UserByRole, error)
	GetRolesByUserID(ctx context.Context, userID string) ([]models.RoleByUser, error)
}

// triển khai Interface IUserRoleService
type UserRoleService struct {
	UserRoleRepo repository.IUserRoleRepository
	logger *zap.Logger
}

func NewUserRoleService(UserRoleRepo repository.IUserRoleRepository, logger *zap.Logger) IUserRoleService {
	return &UserRoleService{
		UserRoleRepo: UserRoleRepo,
		logger: logger,
	}
}

func(s *UserRoleService) Create(ctx context.Context, obj *models.UserRole) (int, error) {
	return s.UserRoleRepo.Create(ctx, obj)
}

func(s *UserRoleService) Update(ctx context.Context, id string, obj *models.UserRole) error {
	return s.UserRoleRepo.Update(ctx, id, obj)
}

func(s *UserRoleService) Delete(ctx context.Context, uuid string, roleID int32) error {
	return s.UserRoleRepo.Delete(ctx, uuid, roleID)
}

func(s *UserRoleService) GetUserByRoleID(ctx context.Context, roleID int32) ([]models.UserByRole, error) {
	return s.UserRoleRepo.GetUserByRoleID(ctx, roleID)
}

func(s *UserRoleService) GetRolesByUserID(ctx context.Context, userID string) ([]models.RoleByUser, error) {
	return s.UserRoleRepo.GetRolesByUserID(ctx, userID)
}
