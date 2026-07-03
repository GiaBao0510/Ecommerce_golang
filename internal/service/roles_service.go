package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
)

// Interface for RolesService
type IRolesService interface {
	GetByID(ctx context.Context, id int32) (*models.Role, error)
	GetAll(ctx context.Context) ([]models.Role, error)
	Create(ctx context.Context, obj *models.Role) (int, error)
	Update(ctx context.Context, id int32, obj *models.Role) error
	Delete(ctx context.Context, id int32) error
}

// Triển khai Interface IRolesService
type RolesService struct {
	RolesRepo repository.IRolesRepository
}

func NewRolesService(repo repository.IRolesRepository) IRolesService {
	return &RolesService{RolesRepo: repo}
}

// CRUD methods for RolesService
func(r *RolesService) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	result, err := r.RolesRepo.GetByID(ctx, id)
	
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(r *RolesService) GetAll(ctx context.Context) ([]models.Role, error) {

	result, err := r.RolesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func(r *RolesService) Create(ctx context.Context, obj *models.Role) (int, error) {

	return r.RolesRepo.Create(ctx, obj)
}

func(r *RolesService) Update(ctx context.Context, id int32, obj *models.Role) error {

	if id <= 0 {
		return apperrors.NewBadRequestError("Mã role không hợp lệ")
	}

	return r.RolesRepo.Update(ctx, id, obj)
}

func(r *RolesService) Delete(ctx context.Context, id int32) error {
	if id <= 0 {
		return apperrors.NewBadRequestError("Mã role không hợp lệ")
	}
	return r.RolesRepo.Delete(ctx, id)
}