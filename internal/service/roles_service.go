package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"go.uber.org/zap"
)

// Interface for RolesService
type IRolesService interface {
	GetByID(ctx context.Context, id int32) (*models.Role, error)
	GetAll(ctx context.Context) ([]models.Role, error)
	Create(ctx context.Context, obj *models.Role) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Role) error
	Update_Patch(ctx context.Context, id int32, obj *models.Role) error
	Delete(ctx context.Context, id int32) error
}

// Triển khai Interface IRolesService
type RolesService struct {
	RolesRepo repository.IRolesRepository
	logger    *zap.Logger
}

func NewRolesService(repo repository.IRolesRepository, logger *zap.Logger) IRolesService {
	return &RolesService{RolesRepo: repo, logger: logger}
}

// CRUD methods for RolesService
func (r *RolesService) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	result, err := r.RolesRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RolesService) GetAll(ctx context.Context) ([]models.Role, error) {

	result, err := r.RolesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (r *RolesService) Create(ctx context.Context, obj *models.Role) (int, error) {

	if obj.Role_name == "" {
		r.logger.Warn("service: Create - Role_name is empty",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Tên role để trống"),
		)
		return 0, apperrors.NewBadRequestError("Tên role không được để trống")
	}

	r.logger.Info("service: Create - Creating new role",
		zap.String("layer:", "Service"),
		zap.String("role_name:", obj.Role_name),
	)
	return r.RolesRepo.Create(ctx, obj)
}

func (r *RolesService) Update_Put(ctx context.Context, id int32, obj *models.Role) error {

	if id <= 0 {
		r.logger.Warn("service: Update_Put - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Mã role không hợp lệ"),
		)
		return apperrors.NewBadRequestError("Mã role không hợp lệ")
	}

	return r.RolesRepo.Update_Put(ctx, id, obj)
}

func (r *RolesService) Update_Patch(ctx context.Context, id int32, obj *models.Role) error {
	if id <= 0 {
		r.logger.Warn("service: Update_Patch - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Mã role không hợp lệ"),
		)
		return apperrors.NewBadRequestError("Mã role không hợp lệ")
	}

	return r.RolesRepo.Update_Patch(ctx, id, obj)
}

func (r *RolesService) Delete(ctx context.Context, id int32) error {
	if id <= 0 {
		r.logger.Warn("service: Delete - Invalid ID",
			zap.String("layer:", "Service"),
			zap.String("reason:", "Mã role không hợp lệ"),
		)
		return apperrors.NewBadRequestError("Mã role không hợp lệ")
	}
	return r.RolesRepo.Delete(ctx, id)
}
