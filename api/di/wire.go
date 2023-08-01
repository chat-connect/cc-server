// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/chat-connect/cc-server/config/database"
    "github.com/chat-connect/cc-server/infra/dao"
    "github.com/chat-connect/cc-server/api/service"	
    "github.com/chat-connect/cc-server/api/presentation/controller"
	"github.com/chat-connect/cc-server/api/presentation/middleware"
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
        service.NewRoomUserService,
        controller.NewRoomUserController,
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
