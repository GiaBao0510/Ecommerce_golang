package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToCategoryModel(category database.Category) models.Category {
	return models.Category{
		IDCate:      category.IDCate,
		Name:        category.Name,
		Description: category.Description.String,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
		DeletedAt:   category.DeletedAt,
	}
}
