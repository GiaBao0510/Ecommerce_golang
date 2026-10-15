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

func InitBrandRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.BrandController, error) {
	wire.Build(
		repositoryimpl.NewBrandRepository,
		service.NewBrandService,
		controller.NewBrandController,
	)
	return new(controller.BrandController), nil
}
