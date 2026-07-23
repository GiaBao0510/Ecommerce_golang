package repositoryimpl

import (
	"context"
	"strconv"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type RolePermissionRepositoryImpl struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

func NewRolePermissionRepositoryImpl(db *database.Queries, logger *zap.Logger) repository.IRolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "RolePermissionRepository"),
	}
}

func (r *RolePermissionRepositoryImpl) GetPermissionsByRoleID(ctx context.Context, id int32) ([]models.Permission, error) {
	rows, err := r.db.GetPermissionsByRoleID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetPermissionsByRoleID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy với ID: " + strconv.Itoa(int(id)))
	}

	if len(rows) == 0 {
		r.dblog.LogWarning("GetPermissionsByRoleID", "No permissions found for role ID: "+strconv.Itoa(int(id)))
		return nil, apperrors.NewNotFoundError("Không tìm thấy quyền nào cho role ID: " + strconv.Itoa(int(id)))
	}

	var result []models.Permission
	for _, row := range rows {
		result = append(result, mapper.ToPermissionByRoleRowModel(row))
	}

	return result, nil
}

func (r *RolePermissionRepositoryImpl) GetRolesByPermissionID(ctx context.Context, id int32) ([]models.Role, error) {
	rows, err := r.db.GetRolesByPermissionID(ctx, id)
	if err != nil {
		r.dblog.LogError("GetRolesByPermissionID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy với ID: " + strconv.Itoa(int(id)))
	}

	if len(rows) == 0 {
		r.dblog.LogWarning("GetRolesByPermissionID", "No roles found for permission ID: "+strconv.Itoa(int(id)))
		return nil, apperrors.NewNotFoundError("Không tìm thấy role nào cho permission ID: " + strconv.Itoa(int(id)))
	}

	var result []models.Role
	for _, row := range rows {
		result = append(result, mapper.ToRoleModel(row))
	}

	return result, nil
}

func (r *RolePermissionRepositoryImpl) Create(ctx context.Context, obj *models.Role_Permission) (int, error) {
	// Điền các giá trị đầu vào
	params := database.CreateRolePermissionParams{
		ActionID: obj.Action_id,
		RoleID:   obj.Role_id,
	}

	// Gọi phương thức CreateRole từ database.Queries và trả về kết quả
	if err := r.db.CreateRolePermission(ctx, params); err != nil {
		r.dblog.LogError("Create", err)
		return 0, apperrors.NewInternalServerError(err)
	}

	return 0, nil
}

func (r *RolePermissionRepositoryImpl) Update_Put(ctx context.Context, obj *models.Role_Permission) error {
	// Điền các giá trị đầu vào
	params := database.UpdateRolePermissionByRoleID_PUTParams{
		ActionID: obj.Action_id,
		RoleID:   obj.Role_id,
	}

	result, err := r.db.UpdateRolePermissionByRoleID_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("Update_Put", err, zap.Int32("action_id", obj.Action_id), zap.Int32("role_id", obj.Role_id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	if affected, err := result.RowsAffected(); err != nil {
		r.dblog.LogError("RowsAffected", err, zap.Int32("action_id", obj.Action_id), zap.Int32("role_id", obj.Role_id))
		return apperrors.NewInternalServerError(err)
	} else if affected == 0 {
		r.dblog.LogWarning("Update_Put", "No rows affected", zap.Int32("action_id", obj.Action_id), zap.Int32("role_id", obj.Role_id))
		return apperrors.NewNotFoundError("Lỗi không tìm thấy role_permission với ID: " + strconv.Itoa(int(obj.Role_id)))
	}

	return nil
}

func (r *RolePermissionRepositoryImpl) Delete(ctx context.Context, role_id, permission_id int32) error {

	// nhận tham số đầu vào
	params := database.DeleteRolePermissionParams{
		ActionID: permission_id,
		RoleID:   role_id,
	}

	result, err := r.db.DeleteRolePermission(ctx, params)
	if err != nil {
		r.dblog.LogError("Delete", err, zap.Int32("role_id", role_id), zap.Int32("permission_id", permission_id))
		return apperrors.NewInternalServerError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		r.dblog.LogError("RowsAffected", err, zap.Int32("role_id", role_id), zap.Int32("permission_id", permission_id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		r.dblog.LogWarning("Delete", "No rows affected", zap.Int32("role_id", role_id), zap.Int32("permission_id", permission_id))
		return apperrors.NewNotFoundError("Lỗi không tìm thấy role_permission với ID: " + strconv.Itoa(int(role_id)))
	}

	return nil
}
