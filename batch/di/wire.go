// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/batch/service"	
    "github.com/game-connect/gc-server/batch/command"
)

// user
func InitializeUserCommand() command.UserCommand {
    wire.Build(
        database.NewDB,
        dao.NewUserDao,
        service.NewUserService,
        command.NewUserCommand,
    )
    return nil
}
