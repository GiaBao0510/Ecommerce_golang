package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)

// Interface for PermissionService
type IPermissionService interface {
	GetByID(ctx context.Context, id int32) (*models.Permission, error)
	GetAll(ctx context.Context) ([]models.Permission, error)
	Create(ctx context.Context, obj *models.Permission) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Permission) error
	Update_Patch(ctx context.Context, id int32, obj *models.Permission) error
	Delete(ctx context.Context, id int32) error
}

// Triển khai các phương thức từ IPermissionService ở đây
type PermissionService struct {
	PermissionRepo repository.IPermissionRepository
	logger         *zap.Logger
}

func NewPermissionService(PermissionRepo repository.IPermissionRepository, logger *zap.Logger) IPermissionService {
	return &PermissionService{PermissionRepo: PermissionRepo, logger: logger}
}

// Triển khai các phương thức từ IPermissionService ở đây
func (s *PermissionService) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	result, err := s.PermissionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *PermissionService) GetAll(ctx context.Context) ([]models.Permission, error) {
	return s.PermissionRepo.GetAll(ctx)

}
func (s *PermissionService) Create(ctx context.Context, obj *models.Permission) (int, error) {

	// Validate input data
	if obj.Action_name == "" {
		s.logger.Warn("service: Create - Action_name is empty",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Tên action để trống"),
		)
		return 0, apperrors.NewBadRequestError("Tên action không được để trống")
	}

	s.logger.Info("service: Create - Creating new permission",
		zap.String("layer:", "Service"),
		zap.String("action_name:", obj.Action_name),
	)
	return s.PermissionRepo.Create(ctx, obj)
}

func (s *PermissionService) Update_Put(ctx context.Context, id int32, obj *models.Permission) error {
	if id <= 0 {
		s.logger.Warn("service: Update_Put - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "ID không hợp lệ"),
		)
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Update_Put(ctx, id, obj)
}

func (s *PermissionService) Update_Patch(ctx context.Context, id int32, obj *models.Permission) error {
	if id <= 0 {
		s.logger.Warn("service: Update_Patch - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "ID không hợp lệ"),
		)
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Update_Patch(ctx, id, obj)
}

func (s *PermissionService) Delete(ctx context.Context, id int32) error {
	if id <= 0 {
		s.logger.Warn("service: Delete - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "ID không hợp lệ"),
		)
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Delete(ctx, id)
}
