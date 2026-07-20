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

// Triển khai các phương thức từ IRolesRepository ở đây
type RolesRepository struct {
	db *database.Queries
	dblog *loghelper.DBLogger
}

// Khởi tạo một RolesRepository mới
func NewRolesRepository(db *database.Queries, logger *zap.Logger) repository.IRolesRepository {
	return &RolesRepository{db: db, dblog: loghelper.NewDBLogger(logger, "RolesRepository")}
}

func (r *RolesRepository) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	rows, err := r.db.GetRoleByID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy role với ID: " + strconv.Itoa(int(id)))
	}

	result := mapper.ToRoleModel(rows)

	return &result, nil
}

func (r *RolesRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	query, err := r.db.GetAllRoles(ctx)

	if err != nil {
		r.dblog.LogError("GetAll", err)
		return nil, err
	}

	if len(query) == 0 {
		r.dblog.LogWarning("GetAll", "No roles found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy role nào")
	}

	var roles []models.Role
	for _, role := range query {
		roles = append(roles, mapper.ToRoleModel(role))
	}

	return roles, nil
}

func (r *RolesRepository) Create(ctx context.Context, obj *models.Role) (int, error) {
	// Điền các giá trị đầu vào
	params := database.CreateRoleParams{
		RoleName: obj.Role_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
	}

	// Gọi phương thức CreateRole từ database.Queries và trả về kết quả
	if err := r.db.CreateRole(ctx, params); err != nil {
		r.dblog.LogError("Create", err)
		return 0, apperrors.NewInternalServerError(err)
	}

	return 0, nil
}

func (r *RolesRepository) Update_Put(ctx context.Context, id int32, obj *models.Role) error {
	// Điền các giá trị đầu vào
	params := database.UpdateRole_PUTParams{
		RoleName: obj.Role_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		RoleID: id,
	}

	result, err := r.db.UpdateRole_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	if affected, err := result.RowsAffected(); err != nil {
		r.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	} else if affected == 0 {
		r.dblog.LogWarning("Update_Put", "No rows affected", zap.Int32("id", id))
		return apperrors.NewNotFoundError("Lỗi không tìm thấy role với ID: " + strconv.Itoa(int(id)))
	}

	return nil
}

func (r *RolesRepository) Update_Patch(ctx context.Context, id int32, obj *models.Role) error {
	// Điền các giá trị đầu vào
	params := database.UpdateRole_PATCHParams{
		RoleName: obj.Role_name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		RoleID: id,
	}

	result, err := r.db.UpdateRole_PATCH(ctx, params)
	if err != nil {
		r.dblog.LogError("Update_Patch", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	if affected, err := result.RowsAffected(); err != nil {
		r.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	} else if affected == 0 {
		r.dblog.LogWarning("Update_Patch", "No rows affected", zap.Int32("id", id))
		return apperrors.NewNotFoundError("Lỗi không tìm thấy role với ID: " + strconv.Itoa(int(id)))
	}

	return nil
}

func (r *RolesRepository) Delete(ctx context.Context, id int32) error {

	result, err := r.db.DeleteRole(ctx, id)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		r.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		r.dblog.LogWarning("Delete", "No rows affected", zap.Int32("id", id))
		return apperrors.NewNotFoundError("Lỗi không tìm thấy role với ID: " + strconv.Itoa(int(id)))
	}

	return nil
}
