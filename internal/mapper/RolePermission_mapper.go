package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToRolePermissionModel(rp database.RolePermission) models.Role_Permission {
	return models.Role_Permission{
		Role_id:       rp.RoleID,
		Permission_id: rp.ActionID,
		Created_at:    rp.CreatedAt,
		Updated_at:    rp.UpdatedAt,
		Deleted_at:    rp.DeletedAt,
	}
}
