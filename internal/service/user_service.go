package service

import (
	"context"
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

// Tạo interface IUserService để định nghĩa các phương thức mà UserService sẽ triển khai
type IUserService interface {
	// CRUD
	GetByID(ctx context.Context, id string) (*models.Users, error)
	Create(ctx context.Context, obj *models.Users) (int, error)
	Update_Put(ctx context.Context, id string, obj *models.Users) error
	Update_Patch(ctx context.Context, id string, obj *models.Users) error
	Delete(ctx context.Context, id string) error

	// List operations 
	GetAll(ctx context.Context) ([]models.Users, error)

	// Search operations
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error)

	// Relationship operations

	// Update other operations
	UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error
}
