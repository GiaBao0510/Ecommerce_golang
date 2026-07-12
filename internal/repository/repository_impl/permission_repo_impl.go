package repositoryimpl

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type PermissionRepository struct {
	db *database.Queries
	dblog *loghelper.DBLogger
}

// Khởi tạo
func NewPermissionRepository(db *database.Queries, logger *zap.Logger) repository.IPermissionRepository {
	return &PermissionRepository{db: db, dblog: loghelper.NewDBLogger(logger, "PermissionRepository")}
}

func(p *PermissionRepository) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	
}

func(p *PermissionRepository) GetAll(ctx context.Context) ([]models.Permission, error)
func(p *PermissionRepository) Create(ctx context.Context, obj *models.Permission) (int, error)
func(p *PermissionRepository) Update_Put(ctx context.Context, id int32, obj *models.Permission) error
func(p *PermissionRepository) Update_Patch(ctx context.Context, id int32, obj *models.Permission) error
func(p *PermissionRepository) Delete(ctx context.Context, id int32) error