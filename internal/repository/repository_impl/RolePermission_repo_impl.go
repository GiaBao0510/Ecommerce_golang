package repositoryimpl

// import (
// 	"context"
// 	"strconv"

// 	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
// 	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
// 	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
// 	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
// 	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
// 	"go.uber.org/zap"
// )

// type RolePermissionRepositoryImpl struct {
// 	db *database.Queries
// 	dblog *loghelper.DBLogger
// }

// func NewRolePermissionRepositoryImpl(db *database.Queries, logger *zap.Logger) repository.IRolePermissionRepository {
// 	return &RolePermissionRepositoryImpl{
// 		db: db,
// 		dblog: loghelper.NewDBLogger(logger, "RolePermissionRepository"),
// 	}
// }

// func toRolePermissionModel(rp database.RolePermission) models.Role_Permission {
// 	return models.Role_Permission{
// 		Role_id: rp.RoleID,
// 		Permission_id: rp.ActionID,
// 		Created_at: rp.CreatedAt,
// 		Updated_at: rp.UpdatedAt,
// 		Deleted_at: rp.DeletedAt,
// 	}
// }

// func(r *RolePermissionRepositoryImpl) GetByID(ctx context.Context, id int32) (*models.Role_Permission, error) {
// 	row, err := r.db.GetPermissionsByRoleID(ctx, id)
// 	if err != nil {
// 		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy với ID: " + strconv.Itoa(int(id)))
// 	}

// 	result := toRolePermissionModel(row)
// 	return &result, nil
// }

// func(r *RolePermissionRepositoryImpl) GetAll(ctx context.Context) ([]models.Role_Permission, error)
// func(r *RolePermissionRepositoryImpl) Create(ctx context.Context, obj *models.Role_Permission) (int, error)
// func(r *RolePermissionRepositoryImpl) Update_Put(ctx context.Context, id int32, obj *models.Role_Permission) error
// func(r *RolePermissionRepositoryImpl) Update_Patch(ctx context.Context, id int32, obj *models.Role_Permission) error
// func(r *RolePermissionRepositoryImpl) Delete(ctx context.Context, id int32) error
