package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type IBrandService interface {
	GetAll(ctx context.Context) ([]models.Brand, error)
	GetByID(ctx context.Context, id int32) (*models.Brand, error)
	GetByName(ctx context.Context, name string) (*models.Brand, error)
	Create(ctx context.Context, obj *models.Brand) error
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.Brand) error
}

type BrandService struct {
	BrandRepo repository.IBrandRepository
	slog      *loghelper.ServiceLogger
}

func NewBrandService(repo repository.IBrandRepository, logger *zap.Logger) IBrandService {
	return &BrandService{
		BrandRepo: repo,
		slog:      loghelper.NewServiceLogger(logger, "BrandService"),
	}
}

func (s *BrandService) GetAll(ctx context.Context) ([]models.Brand, error) {
	return s.BrandRepo.GetAll(ctx)
}

func (s *BrandService) GetByID(ctx context.Context, id int32) (*models.Brand, error) {
	if err := servicesupport.RequirePositiveID32(id, "Brand ID", "get_by_id", s.slog); err != nil {
		return nil, err
	}
	return s.BrandRepo.GetByID(ctx, id)
}

func (s *BrandService) GetByName(ctx context.Context, name string) (*models.Brand, error) {
	if err := servicesupport.RequireNonEmptyString(name, "Brand name", "get_by_name", s.slog); err != nil {
		return nil, err
	}
	return s.BrandRepo.GetByName(ctx, name)
}

func (s *BrandService) Create(ctx context.Context, obj *models.Brand) error {
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "create", s.slog); err != nil {
		return err
	}
	return s.BrandRepo.Create(ctx, obj)
}

func (s *BrandService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Brand ID", "delete", s.slog); err != nil {
		return err
	}
	return s.BrandRepo.Delete(ctx, id)
}

func (s *BrandService) Update_Put(ctx context.Context, id int32, obj *models.Brand) error {
	if err := servicesupport.RequirePositiveID32(id, "Brand ID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "update_put", s.slog); err != nil {
		return err
	}
	return s.BrandRepo.Update_Put(ctx, id, obj)
}
