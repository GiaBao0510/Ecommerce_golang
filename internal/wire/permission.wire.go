//go:build wireinject
// +build wireinject

package wire

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitPermissionRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.PermissionController, error) {
	wire.Build(
		repositoryimpl.NewPermissionRepository,
		service.NewPermissionService,
		controller.NewPermissionController,
	)
	return new(controller.PermissionController), nil
}
