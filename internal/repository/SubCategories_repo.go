package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type ISubCategoryRepository interface {
	GetAll(ctx context.Context) ([]models.SubCategory, error)
	GetByID(ctx context.Context, id int32) (*models.SubCategory, error)
	GetByName(ctx context.Context, name string) (*models.SubCategory, error)
	Create(ctx context.Context, obj *models.SubCategory) (int, error)
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.SubCategory) error
}
