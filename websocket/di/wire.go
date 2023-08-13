// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/websocket/service"	
    "github.com/game-connect/gc-server/websocket/presentation/controller"
	"github.com/game-connect/gc-server/websocket/presentation/middleware"
)

// chat
func InitializeChatController() controller.ChatController {
    wire.Build(
        database.NewDB,
        dao.NewChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewChatService,
        service.NewUserService,
        controller.NewChatController,
    )
    return nil
}

// channel_chat
func InitializeChannelChatController() controller.ChannelChatController {
    wire.Build(
        database.NewDB,
        dao.NewChannelChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewChannelChatService,
        service.NewUserService,
        controller.NewChannelChatController,
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
