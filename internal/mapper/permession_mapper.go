package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

// Tạo Helper để chuyển đổi giữa models.Permission và Database
func ToPermissionModel(p database.Permission) models.Permission {
	return models.Permission{
		Action_id:   p.ActionID,
		Action_name: p.ActionName,
		Description: p.Description.String,
		Created_at:  p.CreatedAt,
		Updated_at:  p.UpdatedAt,
		Deleted_at:  p.UpdatedAt,
	}
}