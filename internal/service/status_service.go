package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
)

// Interface for StatusService
type IStatusService interface {
	GetStatusByID(ctx context.Context, id int32) (*models.Status, error)
	GetAllStatuses(ctx context.Context) ([]models.Status, error)
	CreateStatus(ctx context.Context, obj *models.Status) (int, error)
	UpdateStatus(ctx context.Context, id int32, obj *models.Status) error
	DeleteStatus(ctx context.Context, id int32) error
}

type StatusService struct {
	StatusRepo repository.IStatusRepository
}

// Constructor for StatusService
func NewStatusService(StatusRepo repository.IStatusRepository) IStatusService {
	return &StatusService{
		StatusRepo: StatusRepo,
	}
}

// Triển khai các phương thức từ IStatusService ở đây
func (s *StatusService) GetStatusByID(ctx context.Context, id int32) (*models.Status, error) {
	return s.StatusRepo.GetStatusByID(ctx, id)
}

func (s *StatusService) GetAllStatuses(ctx context.Context) ([]models.Status, error) {
	return s.StatusRepo.GetAllStatuses(ctx)
}

func (s *StatusService) CreateStatus(ctx context.Context, obj *models.Status) (int, error) {
	return s.StatusRepo.CreateStatus(ctx, obj)
}

func (s *StatusService) UpdateStatus(ctx context.Context, id int32, obj *models.Status) error {
	return s.StatusRepo.UpdateStatus(ctx, id, obj)
}

func (s *StatusService) DeleteStatus(ctx context.Context, id int32) error {
	return s.StatusRepo.DeleteStatus(ctx, id)
}

// Các phương thức khác liên quan đến business logic của Status có thể được thêm vào đây
