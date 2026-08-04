// internal/wire/roles.wire.go
//go:build wireinject

// ^ câu lệnh ở trên để áp dụng ràng buộc về xây dựng trong wire, hướng dẫn công cụ xử lý tệp wireinject
package wire

import (
	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http"
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Khởi tạo tiêm phụ thuộc cho RolesController
func InitRolesRouterHandler(db *database.Queries, logger *zap.Logger) (*controller.RolesController, error) {
	wire.Build(
		repositoryimpl.NewRolesRepository, // nhận *database.Queries
		service.NewRolesService,           // nhận IRolesRepository
		controller.NewRolesController,     // nhận IRolesService
	)

	return new(controller.RolesController), nil
}
