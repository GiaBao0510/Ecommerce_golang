package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IProductDetailRepository interface {
	GetAll(ctx context.Context) ([]models.ProductDetail, error)
	GetByID(ctx context.Context, idProduct string) (*models.ProductDetail, error)
	Create(ctx context.Context, obj *models.ProductDetail) error
	Delete(ctx context.Context, idProduct string) error
	Update_Put(ctx context.Context, idProduct string, obj *models.ProductDetail) error
	UpdatePrice_Patch(ctx context.Context, idProduct string, price string) error
	UpdateLikeCount_Patch(ctx context.Context, idProduct string, likeCount int32) error
	UpdateQuantity_Patch(ctx context.Context, idProduct string, quantity int32) error
}
