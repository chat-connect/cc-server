package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/game-connect/gc-server/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-connect/gc-server/log"
	"github.com/game-connect/gc-server/game/di"
)

func Init() {
	// di: wire ./api/di/wire.go
	userController := di.InitializeUserController()
	userMiddleware := di.InitializeUserMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth: 認証API
	auth := e.Group("/auth")
	auth.POST("/register_user", userController.RegisterUser()) // auth/user_register
	auth.POST("/login_user", userController.LoginUser()) // auth/user_login

	// auth: 認証済ユーザーのみアクセス可能
	auth.Use(userMiddleware.UserMiddleware)
	auth.GET("/check_user/:userKey", userController.CheckUser()) // auth/user_check/:userKey
	auth.PUT("/logout_user/:userKey", userController.LogoutUser()) // auth/user_logout/:userKey
	auth.DELETE("/delete_user/:userKey", userController.DeleteUser()) // auth/user_delete/:userKey

	e.Logger.Fatal(e.Start(":8000"))
}
