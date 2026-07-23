package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
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
	logger             *zap.Logger
}

func NewRolePermissionService(repo repository.IRolePermissionRepository, logger *zap.Logger) IRolePermissionService {
	return &RolePermissionService{RolePermissionRepo: repo, logger: logger}
}

func (s *RolePermissionService) GetPermissionsByRoleID(ctx context.Context, id int32) ([]models.Permission, error) {

	result, err := s.RolePermissionRepo.GetPermissionsByRoleID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *RolePermissionService) GetRolesByPermissionID(ctx context.Context, id int32) ([]models.Role, error) {
	result, err := s.RolePermissionRepo.GetRolesByPermissionID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *RolePermissionService) Create(ctx context.Context, obj *models.Role_Permission) (int, error) {
	if obj.Role_id == 0 || obj.Action_id == 0 {
		s.logger.Warn("service: Create - Role_id or Action_id is empty",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Role_id hoặc Action_id để trống"),
		)
		return 0, apperrors.NewBadRequestError("Role_id hoặc Action_id không được để trống")
	}

	s.logger.Info("service: Create - Creating new Role_Permission",
		zap.Int32("role_id", obj.Role_id),
		zap.Int32("action_id", obj.Action_id),
	)

	return s.RolePermissionRepo.Create(ctx, obj)
}

func (s *RolePermissionService) Update_Put(ctx context.Context, obj *models.Role_Permission) error {
	if obj.Role_id == 0 || obj.Action_id == 0 {
		s.logger.Warn("service: Update_Put - Role_id or Action_id is empty",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Role_id hoặc Action_id để trống"),
		)
		return apperrors.NewBadRequestError("Role_id hoặc Action_id không được để trống")
	}

	s.logger.Info("service: Update_Put - Updating Role_Permission",
		zap.Int32("role_id", obj.Role_id),
		zap.Int32("action_id", obj.Action_id),
	)
	return s.RolePermissionRepo.Update_Put(ctx, obj)
}

func (s *RolePermissionService) Delete(ctx context.Context, obj *models.Role_Permission) error {
	if obj.Role_id == 0 || obj.Action_id == 0 {
		s.logger.Warn("service: Delete - Role_id or Action_id is empty",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Role_id hoặc Action_id để trống"),
		)
		return apperrors.NewBadRequestError("Role_id hoặc Action_id không được để trống")
	}
	return s.RolePermissionRepo.Delete(ctx, obj.Role_id, obj.Action_id)
}
