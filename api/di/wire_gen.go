// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/chat-connect/cc-server/api/presentation/controller"
	"github.com/chat-connect/cc-server/api/presentation/middleware"
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/config/database"
	"github.com/chat-connect/cc-server/infra/dao"
)

// Injectors from wire.go:

// user
func InitializeUserController() controller.UserController {
	db := database.NewDB()
	userRepository := dao.NewUserRepository(db)
	transactionRepository := dao.NewTransactionRepository(db)
	userService := service.NewUserService(userRepository, transactionRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// room
func InitializeRoomController() controller.RoomController {
	db := database.NewDB()
	roomRepository := dao.NewRoomRepository(db)
	transactionRepository := dao.NewTransactionRepository(db)
	roomService := service.NewRoomService(roomRepository, transactionRepository)
	roomController := controller.NewRoomController(roomService)
	return roomController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	db := database.NewDB()
	userRepository := dao.NewUserRepository(db)
	transactionRepository := dao.NewTransactionRepository(db)
	userService := service.NewUserService(userRepository, transactionRepository)
	userMiddleware := middleware.NewUserMiddleware(userService)
	return userMiddleware
}
