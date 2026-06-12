package repositoryimpl

import (
	"context"
	"database/sql"
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
)

type StatusRepository struct {
	db *database.Queries
}

// triển khai các phương thứ từ IStatusRepository ở đây
func NewStatusRepository(db *database.Queries) repository.IStatusRepository {
	return &StatusRepository{db: db}
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
func (s *StatusRepository) GetStatusByID(ctx context.Context, id int32) (*models.Status, error) {
	rows, err := s.db.GetStatusByID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := toStatusModel(rows)
	return &result, nil
}

func (s *StatusRepository) GetAllStatuses(ctx context.Context) ([]models.Status, error) {
	query, err := s.db.GetAllStatus(ctx)
	if err != nil{
		return nil, err
	}

	var statuses []models.Status
	for _, status := range query {
		
		statuses = append(statuses, toStatusModel(status))
	}

	return statuses, nil
}

func (s *StatusRepository) CreateStatus(ctx context.Context, obj *models.Status) (int, error) {
	params := database.CreateStatusParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid: obj.Description != "",
		},
		UpdatedAt: obj.Updated_at,
		DeletedAt: obj.Deleted_at,
	}

	// Gọi phương thức CreateStatus từ database.Queries để thực hiện việc tạo mới
	if err := s.db.CreateStatus(ctx, params); err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *StatusRepository) UpdateStatus(ctx context.Context, id int32, obj *models.Status) error {
	params := database.UpdateStatusParams{
		Name: obj.Name,
		Description: sql.NullString{
			String: obj.Description,
			Valid: obj.Description != "",
		},
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},	
	}

	return s.db.UpdateStatus(ctx, params)
}

func (s *StatusRepository) DeleteStatus(ctx context.Context, id int32) error {
	return s.db.DeleteStatus(ctx, id)
}
