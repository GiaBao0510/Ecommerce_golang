package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IShopOwnerRepository interface {
	GetAll(ctx context.Context) ([]models.ShopOwner, error)
	GetByID(ctx context.Context, id int32) (*models.ShopOwner, error)
	GetByName(ctx context.Context, name string) (*models.ShopOwner, error)
	Create(ctx context.Context, obj *models.ShopOwner) error
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.ShopOwner) error
	UpdateName_Patch(ctx context.Context, id int32, name string) error
	UpdateDescription_Patch(ctx context.Context, id int32, description string) error
	UpdateAvatar_Patch(ctx context.Context, id int32, avatarURL string) error
}
