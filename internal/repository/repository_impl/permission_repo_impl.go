package repositoryimpl

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type PermissionRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

// Khởi tạo
func NewPermissionRepository(db *database.Queries, logger *zap.Logger) repository.IPermissionRepository {
	return &PermissionRepository{db: db, dblog: loghelper.NewDBLogger(logger, "PermissionRepository")}
}



func (p *PermissionRepository) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	row, err := p.db.GetPermissionByID(ctx, id)
	if err != nil {
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy với ID: " + strconv.Itoa(int(id)))
	}

	result := mapper.ToPermissionModel(row)
	return &result, nil

}

func (p *PermissionRepository) GetAll(ctx context.Context) ([]models.Permission, error) {
	query, err := p.db.GetAllPermissions(ctx)

	if err != nil {
		p.dblog.LogError("GetAll", err)
		return nil, err
	}

	if len(query) == 0 {
		p.dblog.LogWarning("GetAll", "No permissions found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy permission nào")
	}

	var permissions []models.Permission
	for _, v := range query {
		permissions = append(permissions, mapper.ToPermissionModel(v))
	}

	return permissions, nil
}

func (p *PermissionRepository) Create(ctx context.Context, obj *models.Permission) (int, error) {
	// Điền thông tin vào
	params := database.CreatePermissionParams{
		ActionName: obj.Action_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
	}

	// gọi phương thức tạo
	if err := p.db.CreatePermission(ctx, params); err != nil {
		p.dblog.LogError("Create", err)
		return 0, apperrors.NewInternalServerError(err)
	}

	return 0, nil
}

func (p *PermissionRepository) Update_Put(ctx context.Context, id int32, obj *models.Permission) error {
	// Điên thông tin vào
	params := database.UpdatePermission_PUTParams{
		ActionName: obj.Action_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		ActionID: id,
	}

	result, err := p.db.UpdatePermission_PUT(ctx, params)
	if err != nil {
		p.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra xem có bản ghi nào bị ảnh hưởng không
	if affected, err := result.RowsAffected(); err != nil {
		p.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	} else if affected == 0 {
		p.dblog.LogWarning("Update_Put", "No permission found with ID: "+strconv.Itoa(int(id)), zap.Int32("id", id))
		return apperrors.NewNotFoundError("Permission not found")
	}

	return nil
}

func (p *PermissionRepository) Update_Patch(ctx context.Context, id int32, obj *models.Permission) error {
	// Điên thông tin vào
	params := database.UpdatePermission_PATCHParams{
		ActionName: obj.Action_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		ActionID: id,
	}

	result, err := p.db.UpdatePermission_PATCH(ctx, params)
	if err != nil {
		p.dblog.LogError("Update_Patch", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra xem có bản ghi nào bị ảnh hưởng không
	if affected, err := result.RowsAffected(); err != nil {
		p.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	} else if affected == 0 {
		p.dblog.LogWarning("Update_Patch", "No permission found with ID: "+strconv.Itoa(int(id)), zap.Int32("id", id))
		return apperrors.NewNotFoundError("Permission not found")
	}

	return nil
}

func (p *PermissionRepository) Delete(ctx context.Context, id int32) error {
	result, err := p.db.DeletePermission(ctx, id)
	if err != nil {
		p.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		p.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		p.dblog.LogWarning("Delete", "No permission found with ID: "+strconv.Itoa(int(id)), zap.Int32("id", id))
		return apperrors.NewNotFoundError("Permission not found")
	}

	return nil
}
