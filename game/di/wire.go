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

// admin user
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
func InitializeUserController() controller.UserController {
    wire.Build(
        database.NewDB,
        dao.NewGameUserDao,
        dao.NewTransactionDao,
        service.NewUserService,
        controller.NewUserController,
    )
    return nil
}

// game
func InitializeGameController() controller.GameController {
    wire.Build(
        database.NewDB,
        dao.NewGameDao,
        dao.NewTransactionDao,
        service.NewGameService,
        controller.NewGameController,
    )
    return nil
}

// game score
func InitializeGameScoreController() controller.GameScoreController {
    wire.Build(
        database.NewDB,
        dao.NewGameScoreDao,
        dao.NewTransactionDao,
        service.NewGameScoreService,
        controller.NewGameScoreController,
    )
    return nil
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
    wire.Build(
		middleware.NewUserMiddleware,
    )
    return nil
}

// admin user
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

// game
func InitializeGameMiddleware() middleware.GameMiddleware {
    wire.Build(
        database.NewDB,
        dao.NewGameDao,
        dao.NewTransactionDao,
        service.NewGameService,
        middleware.NewGameMiddleware,
    )
    return nil
}
