package repository

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IUserRoleRepository interface {
	// CRUD
	Create(ctx context.Context, obj *models.UserRole) (int, error)
	Update(ctx context.Context, id string, obj *models.UserRole) error
	Delete(ctx context.Context, uuid string, roleID int32) error
	GetUserByRoleID(ctx context.Context, roleID int32) ([]models.UserByRole, error)
	GetRolesByUserID(ctx context.Context, userID string) ([]models.RoleByUser, error)

	// WithTx: Thực hiện các thao tác trong một giao dịch
	WithTx(tx *sql.Tx) IUserRoleRepository
}