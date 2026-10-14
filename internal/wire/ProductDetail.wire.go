//go:build wireinject

package wire

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitProductDetailRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.ProductDetailController, error) {
	wire.Build(
		repositoryimpl.NewProductDetailRepository,
		service.NewProductDetailService,
		controller.NewProductDetailController,
	)
	return new(controller.ProductDetailController), nil
}
