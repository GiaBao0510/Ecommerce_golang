package repositoryimpl

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
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

// Tạo Hepper để chuyển đổi giữa models.Status và database.Queries trả về
func toStatusModel(s database.Status) models.Status {
	return models.Status{
		Id_status:   s.IDStatus,
		Name:        s.Name,
		Description: s.Description.String,
		Created_at:  s.CreatedAt,
		Updated_at:  s.UpdatedAt,
		Deleted_at:  s.DeletedAt,
	}
}

// Các phương thức CRUD sẽ được triển khai ở đây
func (s *StatusRepository) GetByID(ctx context.Context, id int32) (*models.Status, error) {
	rows, err := s.db.GetStatusByID(ctx, id)
	if err != nil {

		// log lỗi chi tiết
		s.dblog.LogError("GetByID", err, zap.Int32("id", id))
		return nil, apperrors.NewNotFoundError("Lỗi không tìm thấy status với ID: " + strconv.Itoa(int(id)))
	}

	result := toStatusModel(rows)
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
		statuses = append(statuses, toStatusModel(status))
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

func (s *StatusRepository) Update(ctx context.Context, id int32, obj *models.Status) error {

	// Điền các giá trị đầu vào
	params := database.UpdateStatusParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid:  obj.Description != "",
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		IDStatus: id,
	}

	result, err := s.db.UpdateStatus(ctx, params)
	if err != nil {
		s.dblog.LogError("UpdateStatus", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	// Kiểm tra số lượng bản ghi bị ảnh hưởng
	affected, err := result.RowsAffected()
	if err != nil {
		s.dblog.LogError("RowsAffected", err, zap.Int32("id", id))
		return apperrors.NewInternalServerError(err)
	}

	if affected == 0 {
		s.dblog.LogWarning("UpdateStatus", "No rows affected", zap.Int32("id", id))
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
