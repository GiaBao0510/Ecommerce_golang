package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IPermissionRepository interface {
	GetByID(ctx context.Context, id int32) (*models.Permission, error)
	GetAll(ctx context.Context) ([]models.Permission, error)
	Create(ctx context.Context, obj *models.Permission) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Permission) error
	Update_Patch(ctx context.Context, id int32, obj *models.Permission) error
	Delete(ctx context.Context, id int32) error
}