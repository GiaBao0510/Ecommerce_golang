//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	controllerAuth "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/authen"
	controllerEmail "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/email"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	
	"database/sql"
	"github.com/google/wire"
	"go.uber.org/zap"
	"github.com/mailjet/mailjet-apiv3-go"
)

func InitializeAuthService(
	db *sql.DB,
	queries *database.Queries,
	logger *zap.Logger, 
) (*controllerAuth.AuthenController, error) {
	wire.Build(
		NewDBLogger,
		NewMailJectClient,

		// Repositories Layer
		repositoryimpl.NewUserRepository,
		repositoryimpl.NewUserRoleRepository,
		repositoryimpl.NewEmailRepositoryImpl,
		repositoryimpl.NewRedisRepositoryImpl,
		
		// Services Layer
		service.NewRegisterUseCase,
		//service.NewLoginUseCase,
		service.NewVerifyUserUsecase,
		service.NewAuthService,

		//Controller layer
		
		controllerEmail.NewEmailController,
		controllerAuth.NewRegisterController,
		//controllerAuth.NewLoginUseCase,
		controllerAuth.NewAuthenController,
	)

	return new(controllerAuth.AuthenController), nil
}

func NewDBLogger(logger *zap.Logger) *loghelper.DBLogger{
	return loghelper.NewDBLogger(logger, "VerifyFlow")
}

func NewMailJectClient() *mailjet.Client{
	return mailjet.NewMailjetClient(
		global.Config.Authentication.MailJet.API_key,
		global.Config.Authentication.MailJet.Secret_key,
	)
}