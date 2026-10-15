package service

import (
	"context"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/repository"
	servicesupport "github.com/GiaBao0510/Ecommerce_golang/internal/service/service_support"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"go.uber.org/zap"
)

type IShopOwnerService interface {
	GetAll(ctx context.Context) ([]models.ShopOwner, error)
	GetByID(ctx context.Context, id int32) (*models.ShopOwner, error)
	GetByName(ctx context.Context, name string) (*models.ShopOwner, error)
	Create(ctx context.Context, obj *models.ShopOwner) error
	Delete(ctx context.Context, id int32) error
	Update_Put(ctx context.Context, id int32, obj *models.ShopOwner) error
	UpdateName_Patch(ctx context.Context, id int32, name string) error
	UpdateDescription_Patch(ctx context.Context, id int32, description string) error
	UpdateAvatar_Patch(ctx context.Context, id int32, avatarURL string) error
}

type ShopOwnerService struct {
	ShopOwnerRepo repository.IShopOwnerRepository
	slog          *loghelper.ServiceLogger
}

func NewShopOwnerService(repo repository.IShopOwnerRepository, logger *zap.Logger) IShopOwnerService {
	return &ShopOwnerService{
		ShopOwnerRepo: repo,
		slog:          loghelper.NewServiceLogger(logger, "ShopOwnerService"),
	}
}

func (s *ShopOwnerService) GetAll(ctx context.Context) ([]models.ShopOwner, error) {
	return s.ShopOwnerRepo.GetAll(ctx)
}

func (s *ShopOwnerService) GetByID(ctx context.Context, id int32) (*models.ShopOwner, error) {
	if err := servicesupport.RequirePositiveID32(id, "Shop owner ID", "get_by_id", s.slog); err != nil {
		return nil, err
	}
	return s.ShopOwnerRepo.GetByID(ctx, id)
}

func (s *ShopOwnerService) GetByName(ctx context.Context, name string) (*models.ShopOwner, error) {
	if err := servicesupport.RequireNonEmptyString(name, "Shop owner name", "get_by_name", s.slog); err != nil {
		return nil, err
	}
	return s.ShopOwnerRepo.GetByName(ctx, name)
}

func (s *ShopOwnerService) Create(ctx context.Context, obj *models.ShopOwner) error {
	if err := servicesupport.RequireNonEmptyString(obj.Uuid, "UUID", "create", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "create", s.slog); err != nil {
		return err
	}
	return s.ShopOwnerRepo.Create(ctx, obj)
}

func (s *ShopOwnerService) Delete(ctx context.Context, id int32) error {
	if err := servicesupport.RequirePositiveID32(id, "Shop owner ID", "delete", s.slog); err != nil {
		return err
	}
	return s.ShopOwnerRepo.Delete(ctx, id)
}

func (s *ShopOwnerService) Update_Put(ctx context.Context, id int32, obj *models.ShopOwner) error {
	if err := servicesupport.RequirePositiveID32(id, "Shop owner ID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Uuid, "UUID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Name, "Name", "update_put", s.slog); err != nil {
		return err
	}
	return s.ShopOwnerRepo.Update_Put(ctx, id, obj)
}

func (s *ShopOwnerService) UpdateName_Patch(ctx context.Context, id int32, name string) error {
	if err := s.validatePatchID(id, "update_name"); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(name, "Name", "update_name", s.slog); err != nil {
		return err
	}
	return s.ShopOwnerRepo.UpdateName_Patch(ctx, id, name)
}

func (s *ShopOwnerService) UpdateDescription_Patch(ctx context.Context, id int32, description string) error {
	if err := s.validatePatchID(id, "update_description"); err != nil {
		return err
	}
	return s.ShopOwnerRepo.UpdateDescription_Patch(ctx, id, description)
}

func (s *ShopOwnerService) UpdateAvatar_Patch(ctx context.Context, id int32, avatarURL string) error {
	if err := s.validatePatchID(id, "update_avatar"); err != nil {
		return err
	}
	return s.ShopOwnerRepo.UpdateAvatar_Patch(ctx, id, avatarURL)
}

func (s *ShopOwnerService) validatePatchID(id int32, operation string) error {
	if id <= 0 {
		return apperrors.NewBadRequestError("Shop owner ID must be positive for " + operation)
	}
	return nil
}
