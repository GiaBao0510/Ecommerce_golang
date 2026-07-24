package repository

import (
	"context"
	"database/sql"

	dto "github.com/GiaBao0510/Ecommerce_golang/internal/dto/user"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

// Tạo Interface IUserRepo để định nghĩa các phương thức mà UserRepo sẽ triển khai
type IUserRepository interface {
	
	GetByID(ctx context.Context, id string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserByPhone(ctx context.Context, phone sql.NullString) (*models.Users, error)
	GetUID_PasswordHashByEmail(ctx context.Context, email string) (*dto.UserResponseBase, error)
	GetUID_PasswordHashByPhone(ctx context.Context, phone sql.NullString) (*dto.UserResponseBase, error)
	GetAll(ctx context.Context) ([]models.Users, error)
	RegisterUser(ctx context.Context, obj *models.Users) (int, error)
	Create(ctx context.Context, obj *models.Users) (int, error)
	Update_Put(ctx context.Context, id string, obj *models.Users) error
	Update_Patch(ctx context.Context, id string, obj *models.Users) error
	UpdateUserPassword_PATCH(ctx context.Context, id string, passwordHash string) error
	UpdateUserAvatar_PATCH(ctx context.Context, id string, avatarURL string) error
	Delete(ctx context.Context, id string) error

}