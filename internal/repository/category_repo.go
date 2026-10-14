package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type ICategoryRepository interface {
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByID(ctx context.Context, id int32) (*models.Category, error)
	GetByName(ctx context.Context, name string) (*models.Category, error)
	Create(ctx context.Context, obj *models.Category) (int, error)
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.Category) error
}
