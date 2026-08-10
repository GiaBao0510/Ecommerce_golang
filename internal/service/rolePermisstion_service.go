package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type IRolePermissionService interface {
	GetPermissionsByRoleID(ctx context.Context, id int32) ([]models.Permission, error)
	GetRolesByPermissionID(ctx context.Context, id int32) ([]models.Role, error)
	Create(ctx context.Context, obj *models.Role_Permission) (int, error)
	Update_Put(ctx context.Context, obj *models.Role_Permission) error
	Delete(ctx context.Context, obj *models.Role_Permission) error
}

// Triển khai Interface IRolePermissionService
type RolePermissionService struct {
	RolePermissionRepo repository.IRolePermissionRepository
	slog               *loghelper.ServiceLogger
}

func NewRolePermissionService(repo repository.IRolePermissionRepository, logger *zap.Logger) IRolePermissionService {
	return &RolePermissionService{RolePermissionRepo: repo,
		slog: loghelper.NewServiceLogger(logger, "RolePermissionService")}
}

func (s *RolePermissionService) GetPermissionsByRoleID(ctx context.Context, id int32) ([]models.Permission, error) {
	return s.RolePermissionRepo.GetPermissionsByRoleID(ctx, id)
}

func (s *RolePermissionService) GetRolesByPermissionID(ctx context.Context, id int32) ([]models.Role, error) {
	return s.RolePermissionRepo.GetRolesByPermissionID(ctx, id)
}

func (s *RolePermissionService) Create(ctx context.Context, obj *models.Role_Permission) (int, error) {
	if err := servicesupport.RequirePositiveID32(obj.Role_id, "Role ID", "create", s.slog); err != nil {
		return 0, err
	}

	if err := servicesupport.RequirePositiveID32(obj.Action_id, "Action ID", "create", s.slog); err != nil {
		return 0, err
	}

	s.slog.LogInfo("Create", "Tạo Role_Permission mới thành công", zap.Int32("role_id", obj.Role_id), zap.Int32("action_id", obj.Action_id))

	return s.RolePermissionRepo.Create(ctx, obj)
}

func (s *RolePermissionService) Update_Put(ctx context.Context, obj *models.Role_Permission) error {
	if err := servicesupport.RequirePositiveID32(obj.Role_id, "Role ID", "update_put", s.slog); err != nil {
		return err
	}

	if err := servicesupport.RequirePositiveID32(obj.Action_id, "Action ID", "update_put", s.slog); err != nil {
		return err
	}

	s.slog.LogInfo("Update_Put", "Cập nhật Role_Permission thành công", zap.Int32("role_id", obj.Role_id), zap.Int32("action_id", obj.Action_id))
	return s.RolePermissionRepo.Update_Put(ctx, obj)
}

func (s *RolePermissionService) Delete(ctx context.Context, obj *models.Role_Permission) error {
	if err := servicesupport.RequirePositiveID32(obj.Role_id, "Role ID", "delete", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequirePositiveID32(obj.Action_id, "Action ID", "delete", s.slog); err != nil {
		return err
	}
	return s.RolePermissionRepo.Delete(ctx, obj.Role_id, obj.Action_id)
}
