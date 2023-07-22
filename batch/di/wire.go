// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/chat-connect/cc-server/config/database"
    "github.com/chat-connect/cc-server/infra/dao"
    "github.com/chat-connect/cc-server/batch/service"	
    "github.com/chat-connect/cc-server/batch/command"
)

// user
func InitializeUserCommand() command.UserCommand {
    wire.Build(
        database.NewDB,
        dao.NewUserRepository,
        service.NewUserService,
        command.NewUserCommand,
    )
    return nil
}
