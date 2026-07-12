package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
)

// Interface for StatusService
type IStatusService interface {
	GetByID(ctx context.Context, id int32) (*models.Status, error)
	GetAll(ctx context.Context) ([]models.Status, error)
	Create(ctx context.Context, obj *models.Status) (int, error)
	Update_Put(ctx context.Context, id int32, obj *models.Status) error
	Update_Patch(ctx context.Context, id int32, obj *models.Status) error
	Delete(ctx context.Context, id int32) error
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
func (s *StatusService) GetByID(ctx context.Context, id int32) (*models.Status, error) {
	 
	result, err := s.StatusRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StatusService) GetAll(ctx context.Context) ([]models.Status, error) {
	return s.StatusRepo.GetAll(ctx)
}

func (s *StatusService) Create(ctx context.Context, obj *models.Status) (int, error) {
	
	// Validate input data
	if obj.Name == "" {
		return 0, apperrors.NewBadRequestError("Tên trạng thái không được để trống")
	}
	return s.StatusRepo.Create(ctx, obj)
}

func (s *StatusService) Update_Put(ctx context.Context, id int32, obj *models.Status) error {
	if id <= 0 {
		return apperrors.NewBadRequestError("Mã trạng thái không hợp lệ")
	}
	return s.StatusRepo.Update_Put(ctx, id, obj)
}

func (s *StatusService) Update_Patch(ctx context.Context, id int32, obj *models.Status) error {
	if id <= 0 {
		return apperrors.NewBadRequestError("Mã trạng thái không hợp lệ")
	}
	return s.StatusRepo.Update_Patch(ctx, id, obj)
}

func (s *StatusService) Delete(ctx context.Context, id int32) error {
	if id <= 0 {
		return apperrors.NewBadRequestError("Mã trạng thái không hợp lệ")
	}
	return s.StatusRepo.Delete(ctx, id)
}

// Các phương thức khác liên quan đến business logic của Status có thể được thêm vào đây
