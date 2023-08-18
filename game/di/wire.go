// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/game/service"	
    "github.com/game-connect/gc-server/game/presentation/controller"
	"github.com/game-connect/gc-server/game/presentation/middleware"
)

// user
func InitializeAdminUserController() controller.AdminUserController {
    wire.Build(
        database.NewDB,
        dao.NewAdminUserDao,
        dao.NewTransactionDao,
        service.NewAdminUserService,
        controller.NewAdminUserController,
    )
    return nil
}

// user
func InitializeAdminUserMiddleware() middleware.AdminUserMiddleware {
    wire.Build(
        database.NewDB,
        dao.NewAdminUserDao,
        dao.NewTransactionDao,
        service.NewAdminUserService,
		middleware.NewAdminUserMiddleware,
    )
    return nil
}
