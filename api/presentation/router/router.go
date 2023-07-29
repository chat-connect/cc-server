package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/chat-connect/cc-server/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/chat-connect/cc-server/log"
	"github.com/chat-connect/cc-server/api/di"
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
	a := e.Group("/auth")
	a.POST("/user_register", userController.UserRegister()) // auth/user_register
	a.POST("/user_login", userController.UserLogin()) // auth/user_register

	// user: 認証済ユーザーのみアクセス可能
	u := e.Group("/user/:userKey")
	u.Use(userMiddleware.UserMiddleware)
	u.GET("/user_check", userController.UserCheck()) // user/:userKey/user_check
	u.PUT("/user_logout", userController.UserLogout()) // user/:userKey/user_logout
	u.DELETE("/user_delete", userController.UserDelete()) // user/:userKey/user_delete

	e.Logger.Fatal(e.Start(":8000"))
}
