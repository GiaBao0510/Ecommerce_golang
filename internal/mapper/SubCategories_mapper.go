package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToSubCategoryModel(subCategory database.SubCategory) models.SubCategory {
	return models.SubCategory{
		IDSubCate:   subCategory.IDSubCate,
		IDCate:      subCategory.IDCate,
		Name:        subCategory.Name,
		Description: subCategory.Description.String,
		CreatedAt:   subCategory.CreatedAt,
		UpdatedAt:   subCategory.UpdatedAt,
		DeletedAt:   subCategory.DeletedAt,
	}
}
