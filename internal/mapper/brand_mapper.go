package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToBrandModel(brand database.Brand) models.Brand {
	return models.Brand{
		IDBrand:   brand.IDBrand,
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
		DeletedAt: brand.DeletedAt,
	}
}
