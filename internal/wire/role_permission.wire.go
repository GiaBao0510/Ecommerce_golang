// internal/wire/roles.wire.go
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

// KHởi tạo tiêm phụ thuộc cho RolesController
func InitRolePermissionRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.RolePermissionController, error) {
	wire.Build(
		repositoryimpl.NewRolePermissionRepositoryImpl, // nhận *database.Queries
		service.NewRolePermissionService,               // nhận IRolePermissionRepository
		controller.NewRolePermissionController,         // nhận IRolePermissionService
	)

	return new(controller.RolePermissionController), nil
}
