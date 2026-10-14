package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type ISubCategoryService interface {
	GetAll(ctx context.Context) ([]models.SubCategory, error)
	GetByID(ctx context.Context, id int32) (*models.SubCategory, error)
	GetByName(ctx context.Context, name string) (*models.SubCategory, error)
	Create(ctx context.Context, obj *models.SubCategory) (int, error)
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.SubCategory) error
}

type SubCategoryService struct {
	SubCategoryRepo repository.ISubCategoryRepository
	slog            *loghelper.ServiceLogger
}

func NewSubCategoryService(repo repository.ISubCategoryRepository, logger *zap.Logger) ISubCategoryService {
	return &SubCategoryService{
		SubCategoryRepo: repo,
		slog:            loghelper.NewServiceLogger(logger, "SubCategoryService"),
	}
}

func (s *SubCategoryService) GetAll(ctx context.Context) ([]models.SubCategory, error) {
	return s.SubCategoryRepo.GetAll(ctx)
}

func (s *SubCategoryService) GetByID(ctx context.Context, id int32) (*models.SubCategory, error) {
	return s.SubCategoryRepo.GetByID(ctx, id)
}

func (s *SubCategoryService) GetByName(ctx context.Context, name string) (*models.SubCategory, error) {
	if err := servicesupport.RequireNonEmptyString(name, "Subcategory name", "get_by_name", s.slog); err != nil {
		return nil, err
	}
	return s.SubCategoryRepo.GetByName(ctx, name)
}

func (s *SubCategoryService) Create(ctx context.Context, obj *models.SubCategory) (int, error) {
	if err := servicesupport.RequirePositiveID32(obj.IDCate, "Category ID", "create", s.slog); err != nil {
		return 0, err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "create", s.slog); err != nil {
		return 0, err
	}
	return s.SubCategoryRepo.Create(ctx, obj)
}

func (s *SubCategoryService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Subcategory ID", "delete", s.slog); err != nil {
		return err
	}
	return s.SubCategoryRepo.Delete(ctx, id)
}

func (s *SubCategoryService) Update_Put(ctx context.Context, id int32, obj *models.SubCategory) error {
	if err := servicesupport.RequirePositiveID32(id, "Subcategory ID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "update_put", s.slog); err != nil {
		return err
	}
	return s.SubCategoryRepo.Update_Put(ctx, id, obj)
}
