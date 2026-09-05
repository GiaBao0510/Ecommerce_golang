//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	controllerAuth "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/authen"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/google/wire"
	"go.uber.org/zap"
) 

func InitAuthenRouterHandler(
	db *sql.DB,
	queries *database.Queries,
	logger *zap.Logger,
) (controllerAuth.IAuthenController, error) {
	wire.Build(
		NewDBLogger,
		NewMailJectClient,
		NewServiceLogger,

		// Repositories Layer
		repositoryimpl.NewUserRepository,
		repositoryimpl.NewUserRoleRepository,
		repositoryimpl.NewRedisRepositoryImpl,
		repositoryimpl.NewEmailRepositoryImpl,

		// Services Layer
		service.NewVerifyUserUsecase,
		service.NewRegisterUseCase,
		service.NewLoginUseCase,
		service.NewLogoutUseCase,
		service.NewAuthService,

		//Controller layer
		controllerAuth.NewRegisterController,
		controllerAuth.NewLoginController,
		controllerAuth.NewLogoutController,
		controllerAuth.NewAuthenController,
	)

	return nil, nil
}

func NewServiceLogger(logger *zap.Logger) *loghelper.ServiceLogger {
	return loghelper.NewServiceLogger(logger, "AuthFlow")
}
