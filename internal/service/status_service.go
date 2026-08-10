package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
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
	slog    *loghelper.ServiceLogger 
}

// Constructor for StatusService
func NewStatusService(StatusRepo repository.IStatusRepository, logger *zap.Logger) IStatusService {
	return &StatusService{
		StatusRepo: StatusRepo,
		slog: loghelper.NewServiceLogger(logger, "StatusService"),
	}
}

// Triển khai các phương thức từ IStatusService ở đây
func (s *StatusService) GetByID(ctx context.Context, id int32) (*models.Status, error) {
	return s.StatusRepo.GetByID(ctx, id)
}

func (s *StatusService) GetAll(ctx context.Context) ([]models.Status, error) {
	return s.StatusRepo.GetAll(ctx)
}

func (s *StatusService) Create(ctx context.Context, obj *models.Status) (int, error) {
	
	// Validate input data
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "create", s.slog); err != nil {
		return 0, err
	}

	s.slog.LogInfo("Create", "Tạo trạng thái mới thành công", zap.String("status_name", obj.Name))
	return s.StatusRepo.Create(ctx, obj)
}

func (s *StatusService) Update_Put(ctx context.Context, id int32, obj *models.Status) error {
	if err := servicesupport.RequirePositiveID32(id, "Status ID", "update_put", s.slog); err != nil {
		return err
	}
	return s.StatusRepo.Update_Put(ctx, id, obj)
}

func (s *StatusService) Update_Patch(ctx context.Context, id int32, obj *models.Status) error {
	if err := servicesupport.RequirePositiveID32(id, "Status ID", "update_patch", s.slog); err != nil {
		return err
	}
	return s.StatusRepo.Update_Patch(ctx, id, obj)
}

func (s *StatusService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Status ID", "Delete", s.slog); err != nil {
		return err
	}
	return s.StatusRepo.Delete(ctx, id)
}