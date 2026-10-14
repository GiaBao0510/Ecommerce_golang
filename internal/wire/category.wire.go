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

func InitCategoryRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.CategoryController, error) {
	wire.Build(
		repositoryimpl.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
	)
	return new(controller.CategoryController), nil
}
