package repositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/mapper"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type UserRoleRepository struct {
	db    *database.Queries
	dblog *loghelper.DBLogger
}

// triển khai
func NewUserRoleRepository(db *database.Queries, logger *zap.Logger) repository.IUserRoleRepository {
	return &UserRoleRepository{
		db:    db,
		dblog: loghelper.NewDBLogger(logger, "UserRoleRepository"),
	}
}

func (r *UserRoleRepository) Create(ctx context.Context, obj *models.UserRole) (int, error) {
	params := database.CreateUserRoleParams{
		Uuid:   obj.Uuid,
		RoleID: obj.Id_role,
	}

	if err := r.db.CreateUserRole(ctx, params); err != nil {
		r.dblog.LogError("create", err, zap.String("uuid", obj.Uuid), zap.Int32("role_id", obj.Id_role))
		return 0, MapDBErrorWithContext(err, "Lỗi khi tạo user role")
	}
	return 0, nil
}

func (r *UserRoleRepository) Update(ctx context.Context, id string, obj *models.UserRole) error {
	params := database.UpdateUserRoleByUserID_PUTParams{
		Uuid:   obj.Uuid,
		RoleID: obj.Id_role,
	}

	result, err := r.db.UpdateUserRoleByUserID_PUT(ctx, params)
	if err != nil {
		r.dblog.LogError("update", err, zap.String("uuid", obj.Uuid), zap.Int32("role_id", obj.Id_role))
		return MapDBErrorWithContext(err, "Lỗi khi cập nhật user role")
	}

	if err := CheckRowsAffected(
		result,
		"update_user_role",
		"Không tìm thấy user role với uuid: %s và role_id: %d",
		r.dblog,
		zap.String("uuid", obj.Uuid),
		zap.Int32("role_id", obj.Id_role),
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) Delete(ctx context.Context, uuid string, roleID int32) error {
	result, err := r.db.DeleteUserRole(ctx, database.DeleteUserRoleParams{
		Uuid:   uuid,
		RoleID: roleID,
	})

	if err != nil {
		r.dblog.LogError("delete", err, zap.String("uuid", uuid), zap.Int32("role_id", roleID))
		return MapDBErrorWithContext(err, "Lỗi khi xóa user role")
	}

	if err := CheckRowsAffected(
		result,
		"delete_user_role",
		"Không tìm thấy user role với uuid: %s và role_id: %d",
		r.dblog,
		zap.String("uuid", uuid),
		zap.Int32("role_id", roleID),
	); err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) GetUserByRoleID(ctx context.Context, roleID int32) ([]models.UserByRole, error) {
	rows, err := r.db.GetUserByRoleID(ctx, roleID)
	if err != nil {
		r.dblog.LogError("get_user_by_role_id", err, zap.Int32("role_id", roleID))
		return nil, MapDBErrorWithContext(err, "Lỗi khi lấy user theo role_id")
	}

	if len(rows) == 0 {
		r.dblog.LogInfo("get_user_by_role_id", "Không tìm thấy user với role_id: %d", zap.Int32("role_id", roleID))
		return nil, MapDBErrorWithContext(apperrors.NewNotFoundError(fmt.Sprintf("Không tìm thấy role_id: %d", roleID)), fmt.Sprintf("Không tìm thấy role_id: %d", roleID))
	}

	var result []models.UserByRole
	for _, r := range rows {
		result = append(result, mapper.ToUserByRoleModel(r))
	}

	return result, nil
}

func (r *UserRoleRepository) GetRolesByUserID(ctx context.Context, userID string) ([]models.RoleByUser, error) {
	rows, err := r.db.GetRolesByUserID(ctx, userID)
	if err != nil {
		r.dblog.LogError("get_roles_by_user_id", err, zap.String("user_id", userID))
		return nil, MapDBErrorWithContext(err, "Lỗi khi lấy role theo user_id")
	}

	if len(rows) == 0 {
		r.dblog.LogInfo("get_roles_by_user_id", "Không tìm thấy role với user_id: %s", zap.String("user_id", userID))
		return nil, MapDBErrorWithContext(apperrors.NewNotFoundError(fmt.Sprintf("Không tìm thấy user_id: %s", userID)), fmt.Sprintf("Không tìm thấy user_id: %s", userID))
	}

	var result []models.RoleByUser
	for _, r := range rows {
		result = append(result, mapper.ToRoleByUserModel(r))
	}

	return result, nil
}

func (r *UserRoleRepository) WithTx(tx *sql.Tx) repository.IUserRoleRepository{
	return &UserRoleRepository{
		db: r.db.WithTx(tx),
		dblog: r.dblog,
	}
}