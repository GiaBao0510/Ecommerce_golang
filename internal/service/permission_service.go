package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
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
	slog *loghelper.ServiceLogger
}

func NewPermissionService(PermissionRepo repository.IPermissionRepository, logger *zap.Logger) IPermissionService {
	return &PermissionService{PermissionRepo: PermissionRepo, slog: loghelper.NewServiceLogger(logger, "PermissionService")}
}

// Triển khai các phương thức từ IPermissionService ở đây
func (s *PermissionService) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	return s.PermissionRepo.GetByID(ctx, id)
}

func (s *PermissionService) GetAll(ctx context.Context) ([]models.Permission, error) {
	return s.PermissionRepo.GetAll(ctx)

}
func (s *PermissionService) Create(ctx context.Context, obj *models.Permission) (int, error) {

	// Validate input data
	if err := servicesupport.RequireNonEmptyString(obj.Action_name, "Action_name", "create", s.slog); err != nil {
		return 0, err
	}
	
	s.slog.LogInfo("Create", "Tạo permission mới thành công", zap.String("action_name", obj.Action_name))
	return s.PermissionRepo.Create(ctx, obj)
}

func (s *PermissionService) Update_Put(ctx context.Context, id int32, obj *models.Permission) error {
	if servicesupport.RequirePositiveID32(id, "Permission ID", "update_put", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Update_Put(ctx, id, obj)
}

func (s *PermissionService) Update_Patch(ctx context.Context, id int32, obj *models.Permission) error {
	if servicesupport.RequirePositiveID32(id, "Permission ID", "update_patch", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Update_Patch(ctx, id, obj)
}

func (s *PermissionService) Delete(ctx context.Context, id int32) error {
	if servicesupport.RequirePositiveID32(id, "Permission ID", "delete", s.slog) != nil {
		return apperrors.NewBadRequestError("ID không hợp lệ")
	}
	return s.PermissionRepo.Delete(ctx, id)
}
