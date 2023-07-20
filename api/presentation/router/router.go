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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GeneratApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth: 認証API
	a := e.Group("/auth")
	a.POST("/user_register", userController.UserRegister()) // auth/user_register
	a.POST("/user_login", userController.UserLogin()) // auth/user_register

		// user: 認証済ユーザーのみアクセス可能
	u := e.Group("/user")
	u.Use(userMiddleware.UserMiddleware)
	u.GET("/:userKey/user_check", userController.UserCheck()) // user/user_check
	u.PUT("/:userKey/user_logout", userController.UserLogout()) // user/user_logout
	u.DELETE("/:userKey/user_delete", userController.UserDelete()) // user/user_delete

	e.Logger.Fatal(e.Start(":8000"))
}
