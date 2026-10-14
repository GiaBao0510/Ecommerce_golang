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

func InitSubCategoryRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.SubCategoryController, error) {
	wire.Build(
		repositoryimpl.NewSubCategoryRepository,
		service.NewSubCategoryService,
		controller.NewSubCategoryController,
	)
	return new(controller.SubCategoryController), nil
}
