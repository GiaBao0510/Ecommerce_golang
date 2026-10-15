package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IBrandRepository interface {
	GetAll(ctx context.Context) ([]models.Brand, error)
	GetByID(ctx context.Context, id int32) (*models.Brand, error)
	GetByName(ctx context.Context, name string) (*models.Brand, error)
	Create(ctx context.Context, obj *models.Brand) error
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.Brand) error
}
