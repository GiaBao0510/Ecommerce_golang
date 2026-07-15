package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IRolePermissionRepository interface {
	GetByID(ctx context.Context, id int32) (*models.Role_Permission, error)
	GetAll(ctx context.Context) ([]models.Role_Permission, error)
	Create(ctx context.Context, obj *models.Role_Permission) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Role_Permission) error
	Update_Patch(ctx context.Context, id int32, obj *models.Role_Permission) error
	Delete(ctx context.Context, id int32) error
}