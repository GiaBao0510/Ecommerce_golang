package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToShopOwnerModel(shopOwner database.ShopOwner) models.ShopOwner {
	return models.ShopOwner{
		IDShopOwner: shopOwner.IDShopOwner,
		Uuid:        shopOwner.Uuid,
		Name:        shopOwner.Name,
		Description: shopOwner.Description.String,
		AvatarUrl:   shopOwner.AvatarUrl.String,
		CreatedAt:   shopOwner.CreatedAt,
		UpdatedAt:   shopOwner.UpdatedAt,
		DeletedAt:   shopOwner.DeletedAt,
	}
}
