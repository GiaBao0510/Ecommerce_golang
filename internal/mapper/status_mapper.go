package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToStatusModel(s database.Status) models.Status {
	return models.Status{
		Id_status:   s.IDStatus,
		Name:        s.Name,
		Description: s.Description.String,
		Created_at:  s.CreatedAt,
		Updated_at:  s.UpdatedAt,
		Deleted_at:  s.DeletedAt,
	}
}