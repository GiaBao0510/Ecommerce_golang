package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type ICategoryService interface {
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByID(ctx context.Context, id int32) (*models.Category, error)
	GetByName(ctx context.Context, name string) (*models.Category, error)
	Create(ctx context.Context, obj *models.Category) (int, error)
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.Category) error
}

type CategoryService struct {
	CategoryRepo repository.ICategoryRepository
	slog         *loghelper.ServiceLogger
}

func NewCategoryService(repo repository.ICategoryRepository, logger *zap.Logger) ICategoryService {
	return &CategoryService{
		CategoryRepo: repo,
		slog:         loghelper.NewServiceLogger(logger, "CategoryService"),
	}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.CategoryRepo.GetAll(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, id int32) (*models.Category, error) {
	return s.CategoryRepo.GetByID(ctx, id)
}

func (s *CategoryService) GetByName(ctx context.Context, name string) (*models.Category, error) {
	if err := servicesupport.RequireNonEmptyString(name, "Category name", "get_by_name", s.slog); err != nil {
		return nil, err
	}
	return s.CategoryRepo.GetByName(ctx, name)
}

func (s *CategoryService) Create(ctx context.Context, obj *models.Category) (int, error) {
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "create", s.slog); err != nil {
		return 0, err
	}
	return s.CategoryRepo.Create(ctx, obj)
}

func (s *CategoryService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Category ID", "delete", s.slog); err != nil {
		return err
	}
	return s.CategoryRepo.Delete(ctx, id)
}

func (s *CategoryService) Update_Put(ctx context.Context, id int32, obj *models.Category) error {
	if err := servicesupport.RequirePositiveID32(id, "Category ID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "update_put", s.slog); err != nil {
		return err
	}
	return s.CategoryRepo.Update_Put(ctx, id, obj)
}
