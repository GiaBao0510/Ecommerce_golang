package user

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
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
	slog *loghelper.ServiceLogger
}

func NewUserRoleService(UserRoleRepo repository.IUserRoleRepository, logger *zap.Logger) IUserRoleService {
	return &UserRoleService{
		UserRoleRepo: UserRoleRepo,
		slog: loghelper.NewServiceLogger(logger, "UserRoleService"),
	}
}

func(s *UserRoleService) Create(ctx context.Context, obj *models.UserRole) (int, error) {
	
	if err := servicesupport.RequireNonEmptyString(obj.Uuid, "User UUID", "create", s.slog); err != nil {
		return 0, err
	}

	if err := servicesupport.RequirePositiveID32(obj.Id_role, "Role ID", "create", s.slog); err != nil {
		return 0, err
	}
	
	return s.UserRoleRepo.Create(ctx, obj)
}

func(s *UserRoleService) Update(ctx context.Context, id string, obj *models.UserRole) error {
	
	if err := servicesupport.RequireNonEmptyString(id, "User UUID", "update", s.slog); err != nil {
		return err
	}

	if err := servicesupport.RequirePositiveID32(obj.Id_role, "Role ID", "update", s.slog); err != nil {
		return err
	}
	
	return s.UserRoleRepo.Update(ctx, id, obj)
}

func(s *UserRoleService) Delete(ctx context.Context, uuid string, roleID int32) error {
	
	if err := servicesupport.RequireNonEmptyString(uuid, "User UUID", "delete", s.slog); err != nil {
		return err
	}

	if err := servicesupport.RequirePositiveID32(roleID, "Role ID", "delete", s.slog); err != nil {
		return err
	}
	return s.UserRoleRepo.Delete(ctx, uuid, roleID)
}

func(s *UserRoleService) GetUserByRoleID(ctx context.Context, roleID int32) ([]models.UserByRole, error) {
	return s.UserRoleRepo.GetUserByRoleID(ctx, roleID)
}

func(s *UserRoleService) GetRolesByUserID(ctx context.Context, userID string) ([]models.RoleByUser, error) {
	return s.UserRoleRepo.GetRolesByUserID(ctx, userID)
}
