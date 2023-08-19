// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-connect/gc-server/config/database"
	"github.com/game-connect/gc-server/game/presentation/controller"
	"github.com/game-connect/gc-server/game/presentation/middleware"
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/infra/dao"
)

// Injectors from wire.go:

// admin user
func InitializeAdminUserController() controller.AdminUserController {
	db := database.NewDB()
	adminUserRepository := dao.NewAdminUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	adminUserService := service.NewAdminUserService(adminUserRepository, transactionRepository)
	adminUserController := controller.NewAdminUserController(adminUserService)
	return adminUserController
}

// user
func InitializeUserController() controller.UserController {
	db := database.NewDB()
	gameUserRepository := dao.NewGameUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(gameUserRepository, transactionRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// game
func InitializeGameController() controller.GameController {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	genreRepository := dao.NewGenreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameService := service.NewGameService(gameRepository, genreRepository, transactionRepository)
	gameController := controller.NewGameController(gameService)
	return gameController
}

// game score
func InitializeGameScoreController() controller.GameScoreController {
	db := database.NewDB()
	gameScoreRepository := dao.NewGameScoreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameScoreService := service.NewGameScoreService(gameScoreRepository, transactionRepository)
	gameScoreController := controller.NewGameScoreController(gameScoreService)
	return gameScoreController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	userMiddleware := middleware.NewUserMiddleware()
	return userMiddleware
}

// admin user
func InitializeAdminUserMiddleware() middleware.AdminUserMiddleware {
	db := database.NewDB()
	adminUserRepository := dao.NewAdminUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	adminUserService := service.NewAdminUserService(adminUserRepository, transactionRepository)
	adminUserMiddleware := middleware.NewAdminUserMiddleware(adminUserService)
	return adminUserMiddleware
}

// game
func InitializeGameMiddleware() middleware.GameMiddleware {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	genreRepository := dao.NewGenreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameService := service.NewGameService(gameRepository, genreRepository, transactionRepository)
	gameMiddleware := middleware.NewGameMiddleware(gameService)
	return gameMiddleware
}
