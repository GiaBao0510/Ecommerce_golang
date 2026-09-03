//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/GiaBao0510/Ecommerce_golang/global"
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/email"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
	"github.com/google/wire"
	"github.com/mailjet/mailjet-apiv3-go"
	"go.uber.org/zap"
)

func InitVerifyRouterHandler(
	db *database.Queries,
	logger *zap.Logger,
) (controller.EmailControllerInterface, error) {
	wire.Build(
		NewDBLogger,
		NewMailJectClient,

		repositoryimpl.NewUserRepository,
		repositoryimpl.NewEmailRepositoryImpl,
		repositoryimpl.NewRedisRepositoryImpl,

		service.NewVerifyUserUsecase,
		controller.NewEmailController,
	)

	return nil, nil
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