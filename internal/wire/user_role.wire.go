//go:build wireinject

package wire

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/user"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitUserRoleRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.UserRoleController, error) {
	wire.Build(
		repositoryimpl.NewUserRoleRepository,
		service.NewUserRoleService,
		controller.NewUserRoleController,
	)

	return new(controller.UserRoleController), nil
}
