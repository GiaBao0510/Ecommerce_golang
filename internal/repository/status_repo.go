package repository

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

type IStatusRepository interface {
	//CRUD
	GetStatusByID(ctx context.Context, id int32) (*models.Status, error)
	GetAllStatuses(ctx context.Context) ([]models.Status, error)
	CreateStatus(ctx context.Context, obj *models.Status) (int, error)
	UpdateStatus(ctx context.Context, id int32, obj *models.Status) error
	DeleteStatus(ctx context.Context, id int32) error
}