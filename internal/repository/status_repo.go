package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IStatusRepository interface {
	//CRUD
	GetByID(ctx context.Context, id int32) (*models.Status, error)
	GetAll(ctx context.Context) ([]models.Status, error)
	Create(ctx context.Context, obj *models.Status) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Status) error
	Update_Patch(ctx context.Context, id int32, obj *models.Status) error
	Delete(ctx context.Context, id int32) error
}
