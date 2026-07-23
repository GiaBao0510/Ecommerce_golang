package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IRolePermissionRepository interface {
	GetPermissionsByRoleID(ctx context.Context, id int32) ([]models.Permission, error)
	GetRolesByPermissionID(ctx context.Context, id int32) ([]models.Role, error)
	Create(ctx context.Context, obj *models.Role_Permission) (int, error)
	Update_Put(ctx context.Context, obj *models.Role_Permission) error
	Delete(ctx context.Context, role_id, permission_id int32) error
}
