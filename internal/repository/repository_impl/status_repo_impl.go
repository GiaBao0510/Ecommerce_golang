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

type StatusRepository struct {
	db     *database.Queries
	dblog *loghelper.DBLogger
}

// triển khai các phương thứ từ IStatusRepository ở đây
func NewStatusRepository(db *database.Queries, logger *zap.Logger) repository.IStatusRepository {
	return &StatusRepository{db: db, dblog: loghelper.NewDBLogger(logger, "StatusRepository")}
}

// Các phương thức CRUD sẽ được triển khai ở đây
func (s *StatusRepository) GetByID(ctx context.Context, id int32) (*models.Status, error) {
	rows, err := s.db.GetStatusByID(ctx, id)
	if err != nil {

		// log lỗi chi tiết
		s.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy status với ID: " + strconv.Itoa(int(id)))
	}

	result := mapper.ToStatusModel(rows)
	return &result, nil
}

func (s *StatusRepository) GetAll(ctx context.Context) ([]models.Status, error) {
	query, err := s.db.GetAllStatus(ctx)
	if err != nil {
		s.dblog.LogError("GetAll", err)
		return nil, err
	}

	if len(query) == 0 {
		s.dblog.LogWarning("GetAll", "No statuses found")
		return nil, apperrors.NewNotFoundError("Không tìm thấy status nào")
	}

	var statuses []models.Status
	for _, status := range query {
		statuses = append(statuses, mapper.ToStatusModel(status))
	}

	return statuses, nil
}

func (s *StatusRepository) Create(ctx context.Context, obj *models.Status) (int, error) {
	params := database.CreateStatusParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		UpdatedAt: obj.Updated_at,
		DeletedAt: obj.Deleted_at,
	}

	// Gọi phương thức CreateStatus từ database.Queries để thực hiện việc tạo mới
	if err := s.db.CreateStatus(ctx, params); err != nil {
		s.dblog.LogError("Create", err, zap.String("name", obj.Name))
		return 0, apperrors.NewInternalServerError(err)
	}

	return 0, nil
}

func (s *StatusRepository) Update_Put(ctx context.Context, id int32, obj *models.Status) error {

	// Điền các giá trị đầu vào
	params := database.UpdateStatus_PUTParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		IDStatus: id,
	}

	result, err := s.db.UpdateStatus_PUT(ctx, params)
	if err != nil {
		s.dblog.LogError("UpdateStatus_PUT", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	if err := CheckRowsAffected(
		result,
		"UpdateStatus_PUT",
		"Không tìm thấy status với ID: "+strconv.Itoa(int(id)),
		s.dblog,
		zap.Int32("id", id),
	); err != nil {
		return err
	}

	return nil
}

func (s *StatusRepository) Update_Patch(ctx context.Context, id int32, obj *models.Status) error {

	// Điền các giá trị đầu vào
	params := database.UpdateStatus_PATCHParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		IDStatus: id,
	}

	result, err := s.db.UpdateStatus_PATCH(ctx, params)
	if err != nil {
		s.dblog.LogError("UpdateStatus_PATCH", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	affected, err := result.RowsAffected()
	if err != nil {
		s.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		s.dblog.LogWarning("UpdateStatus_PATCH", "No rows affected", zap.Int32("id", id))
		return apperrors.NewNotFoundError("Không tìm thấy status với ID: " + strconv.Itoa(int(id)))
	}

	return nil
}

func (s *StatusRepository) Delete(ctx context.Context, id int32) error {

	result, err := s.db.DeleteStatus(ctx, id)
	if err != nil {
		s.dblog.LogError("Delete", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		s.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		s.dblog.LogWarning("Delete", "No rows affected", zap.Int32("id", id))
		return apperrors.NewNotFoundError("Không tìm thấy status với ID: " + strconv.Itoa(int(id)))
	}

	return nil
}
