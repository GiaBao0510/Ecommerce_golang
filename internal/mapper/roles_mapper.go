package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToRoleModel(r database.Role) models.Role {
	return models.Role{
		Role_id:     r.RoleID,
		Role_name:   r.RoleName,
		Description: r.Description.String,
	}
}