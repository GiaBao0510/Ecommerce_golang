package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToPermissionByRoleRowModel(r database.GetPermissionsByRoleIDRow) models.Permission {
	return models.Permission{
		Action_id:   r.ActionID,
		Action_name: r.ActionName,
		Description: r.Description.String,
	}
}