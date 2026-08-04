// internal/wire/status.wire.go
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

func InitStatusRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.StatusController, error) {
	wire.Build(
		repositoryimpl.NewStatusRepository, // nhận *database.Queries
		service.NewStatusService,           // nhận IStatusRepository
		controller.NewStatusController,     // nhận IStatusService
	)
	return new(controller.StatusController), nil
}
