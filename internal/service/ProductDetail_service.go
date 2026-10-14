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

type IProductDetailService interface {
	GetAll(ctx context.Context) ([]models.ProductDetail, error)
	GetByID(ctx context.Context, idProduct string) (*models.ProductDetail, error)
	Create(ctx context.Context, obj *models.ProductDetail) error
	Delete(ctx context.Context, idProduct string) error
	Update_Put(ctx context.Context, idProduct string, obj *models.ProductDetail) error
	UpdatePrice_Patch(ctx context.Context, idProduct string, price string) error
	UpdateLikeCount_Patch(ctx context.Context, idProduct string, likeCount int32) error
	UpdateQuantity_Patch(ctx context.Context, idProduct string, quantity int32) error
}

type ProductDetailService struct {
	ProductDetailRepo repository.IProductDetailRepository
	slog              *loghelper.ServiceLogger
}

func NewProductDetailService(repo repository.IProductDetailRepository, logger *zap.Logger) IProductDetailService {
	return &ProductDetailService{
		ProductDetailRepo: repo,
		slog:              loghelper.NewServiceLogger(logger, "ProductDetailService"),
	}
}

func (s *ProductDetailService) GetAll(ctx context.Context) ([]models.ProductDetail, error) {
	return s.ProductDetailRepo.GetAll(ctx)
}

func (s *ProductDetailService) GetByID(ctx context.Context, idProduct string) (*models.ProductDetail, error) {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "get_by_id", s.slog); err != nil {
		return nil, err
	}
	return s.ProductDetailRepo.GetByID(ctx, idProduct)
}

func (s *ProductDetailService) Create(ctx context.Context, obj *models.ProductDetail) error {
	if err := servicesupport.RequireNonEmptyString(obj.IDProduct, "Product ID", "create", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Price, "Price", "create", s.slog); err != nil {
		return err
	}
	return s.ProductDetailRepo.Create(ctx, obj)
}

func (s *ProductDetailService) Delete(ctx context.Context, idProduct string) error {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "delete", s.slog); err != nil {
		return err
	}
	return s.ProductDetailRepo.Delete(ctx, idProduct)
}

func (s *ProductDetailService) Update_Put(ctx context.Context, idProduct string, obj *models.ProductDetail) error {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "update_put", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(obj.Price, "Price", "update_put", s.slog); err != nil {
		return err
	}
	return s.ProductDetailRepo.Update_Put(ctx, idProduct, obj)
}

func (s *ProductDetailService) UpdatePrice_Patch(ctx context.Context, idProduct string, price string) error {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "update_price", s.slog); err != nil {
		return err
	}
	if err := servicesupport.RequireNonEmptyString(price, "Price", "update_price", s.slog); err != nil {
		return err
	}
	return s.ProductDetailRepo.UpdatePrice_Patch(ctx, idProduct, price)
}

func (s *ProductDetailService) UpdateLikeCount_Patch(ctx context.Context, idProduct string, likeCount int32) error {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "update_like_count", s.slog); err != nil {
		return err
	}
	if likeCount < 0 {
		return apperrors.NewBadRequestError("Like count cannot be negative")
	}
	return s.ProductDetailRepo.UpdateLikeCount_Patch(ctx, idProduct, likeCount)
}

func (s *ProductDetailService) UpdateQuantity_Patch(ctx context.Context, idProduct string, quantity int32) error {
	if err := servicesupport.RequireNonEmptyString(idProduct, "Product ID", "update_quantity", s.slog); err != nil {
		return err
	}
	if quantity < 0 {
		return apperrors.NewBadRequestError("Quantity cannot be negative")
	}
	return s.ProductDetailRepo.UpdateQuantity_Patch(ctx, idProduct, quantity)
}
