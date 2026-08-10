package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
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
	slog    *loghelper.ServiceLogger
}

func NewRolesService(repo repository.IRolesRepository, logger *zap.Logger) IRolesService {
	return &RolesService{
		RolesRepo: repo, 
		slog: loghelper.NewServiceLogger(logger, "RolesService"),
	}
}

// CRUD methods for RolesService
func (r *RolesService) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	return r.RolesRepo.GetByID(ctx, id)
}

func (r *RolesService) GetAll(ctx context.Context) ([]models.Role, error) {
	return r.RolesRepo.GetAll(ctx)
}

func (r *RolesService) Create(ctx context.Context, obj *models.Role) (int, error) {

	if err := servicesupport.RequireNonEmptyString(obj.Role_name, "Role_name", "create", r.slog); err != nil {
		return 0, err
	}

	r.slog.LogInfo("Create", "Tạo role mới thành công", zap.String("role_name", obj.Role_name))
	return r.RolesRepo.Create(ctx, obj)
}

func (r *RolesService) Update_Put(ctx context.Context, id int32, obj *models.Role) error {

	if err := servicesupport.RequirePositiveID32(id, "Role ID", "updat_put", r.slog); err != nil {
		return err
	}

	return r.RolesRepo.Update_Put(ctx, id, obj)
}

func (r *RolesService) Update_Patch(ctx context.Context, id int32, obj *models.Role) error {
	if err := servicesupport.RequirePositiveID32(id, "Role ID", "updat_patch", r.slog); err != nil {
		return err
	}

	return r.RolesRepo.Update_Patch(ctx, id, obj)
}

func (r *RolesService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Mã role", "Delete", r.slog); err != nil {
		return err
	}
	return r.RolesRepo.Delete(ctx, id)
}
