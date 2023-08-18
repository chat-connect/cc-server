// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/admin/service"	
    "github.com/game-connect/gc-server/admin/presentation/controller"
	"github.com/game-connect/gc-server/admin/presentation/middleware"
)

// user
func InitializeUserController() controller.UserController {
    wire.Build(
        database.NewDB,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewUserService,
        controller.NewUserController,
    )
    return nil
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
    wire.Build(
        database.NewDB,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewUserService,
		middleware.NewUserMiddleware,
    )
    return nil
}
