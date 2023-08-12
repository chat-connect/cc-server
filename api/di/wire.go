// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/api/service"	
    "github.com/game-connect/gc-server/api/presentation/controller"
	"github.com/game-connect/gc-server/api/presentation/middleware"
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

// room
func InitializeRoomController() controller.RoomController {
    wire.Build(
        database.NewDB,
        dao.NewRoomDao,
        dao.NewRoomUserDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewRoomService,
        controller.NewRoomController,
    )
    return nil
}

// room_user
func InitializeRoomUserController() controller.RoomUserController {
    wire.Build(
        database.NewDB,
        dao.NewRoomDao,
        dao.NewRoomUserDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewRoomService,
        service.NewRoomUserService,
        controller.NewRoomUserController,
    )
    return nil
}

// channel
func InitializeChannelController() controller.ChannelController {
    wire.Build(
        database.NewDB,
        dao.NewChannelDao,
        dao.NewChatDao,
        dao.NewTransactionDao,
        service.NewChannelService,
        controller.NewChannelController,
    )
    return nil
}

// chat
func InitializeChatController() controller.ChatController {
    wire.Build(
        database.NewDB,
        dao.NewChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewChatService,
        controller.NewChatController,
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
