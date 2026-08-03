//go:build wireinject

package wire

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/controller"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitUserRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.UserController, error) {
	wire.Build(
		repositoryimpl.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
