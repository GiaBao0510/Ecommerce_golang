//go:build wireinject
// +build wireinject

package wire

// import (
// 	"database/sql"

// 	controller "github.com/GiaBao0510/Ecommerce_golang/internal/controller/http/authen"
// 	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
// 	repositoryimpl "github.com/GiaBao0510/Ecommerce_golang/internal/repository/repository_impl"
// 	service "github.com/GiaBao0510/Ecommerce_golang/internal/service/authen"
// 	"github.com/google/wire"
// 	"go.uber.org/zap"
// )

// func InitAuthenRouterHandler(db *sql.DB,queries *database.Queries,logger *zap.Logger) (*controller., error) {
// 	wire.Build( 
// 		db,
// 		queries,
// 		repositoryimpl.NewUserRepository,
// 		repositoryimpl.NewUserRoleRepository,
// 		repositoryimpl.NewRedisRepositoryImpl,
// 		service.NewRegisterUseCase,
// 		service.NewLoginUseCase,
// 		service.NewVerifyUserUsecase,
// 		controller.NewRegisterController,
// 		controller.NewLoginController,
// 		controller.NewLogoutController,
// 		controller.NewAuthenController,
// 	)

// 	return new(controller.AuthenController), nil
// }