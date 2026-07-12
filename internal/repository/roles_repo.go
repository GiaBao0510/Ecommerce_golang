package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

// CRUD interface for Roles
type IRolesRepository interface {
	GetByID(ctx context.Context, id int32) (*models.Role, error)
	GetAll(ctx context.Context) ([]models.Role, error)
	Create(ctx context.Context, obj *models.Role) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Role) error
	Update_Patch(ctx context.Context, id int32, obj *models.Role) error
	Delete(ctx context.Context, id int32) error
}