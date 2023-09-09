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
        dao.NewFollowDao,
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
        dao.NewFollowDao,
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

// follow
func InitializeFollowController() controller.FollowController {
    wire.Build(
        database.NewDB,
        dao.NewFollowDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewFollowService,
        controller.NewFollowController,
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

// room_chat
func InitializeRoomChatController() controller.RoomChatController {
    wire.Build(
        database.NewDB,
        dao.NewRoomChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewRoomChatService,
        controller.NewRoomChatController,
    )
    return nil
}

// open_chat
func InitializeOpenChatController() controller.OpenChatController {
    wire.Build(
        database.NewDB,
        dao.NewOpenChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewOpenChatService,
        controller.NewOpenChatController,
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
        controller.NewChannelChatController,
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
