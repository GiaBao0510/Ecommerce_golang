// internal/wire/status.wire.go
//go:build wireinject

package wire

import (
    "github.com/GiaBao0510/Ecommerce_golang/internal/controller"
    "github.com/GiaBao0510/Ecommerce_golang/internal/database"
    repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
    "github.com/GiaBao0510/Ecommerce_golang/internal/service"
    "github.com/google/wire"
)

func InitStatusRouterHandler(db *database.Queries) (*controller.StatusController, error) {
    wire.Build(
        repositoryimpl.NewStatusRepository, // nhận *database.Queries
        service.NewStatusService,           // nhận IStatusRepository
        controller.NewStatusController,     // nhận IStatusService
    )
    return new(controller.StatusController), nil
}