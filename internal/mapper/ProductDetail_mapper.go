package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToProductDetailModel(productDetail database.ProductDetail) models.ProductDetail {
	return models.ProductDetail{
		IDProduct:      productDetail.IDProduct,
		Price:          productDetail.Price,
		LikeCount:      productDetail.LikeCount,
		Quantity:       productDetail.Quantity,
		ProductionDate: productDetail.ProductionDate,
		ExpirationDate: productDetail.ExpirationDate,
		UpdatedAt:      productDetail.UpdatedAt,
	}
}
