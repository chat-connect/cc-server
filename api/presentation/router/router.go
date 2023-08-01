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
	roomController := di.InitializeRoomController()
	roomUserController := di.InitializeRoomUserController()
	userMiddleware := di.InitializeUserMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth: 認証API
	a := e.Group("/auth")
	a.POST("/user_register", userController.UserRegister()) // auth/user_register
	a.POST("/user_login", userController.UserLogin()) // auth/user_login

	// auth: 認証済ユーザーのみアクセス可能
	a.Use(userMiddleware.UserMiddleware)
	a.GET("/user_check/:userKey", userController.UserCheck()) // auth/user_check/:userKey
	a.PUT("/user_logout/:userKey", userController.UserLogout()) // auth/user_logout/:userKey
	a.DELETE("/user_delete/:userKey", userController.UserDelete()) // auth/user_delete/:userKey

	// room: 部屋関連
	r := e.Group("/room")
	r.POST("/:userKey/room_create", roomController.RoomCreate()) // room/:userKey/room_create
	r.POST("/:userKey/room_join/:roomKey", roomUserController.RoomJoin()) // room/:userKey/room_join/:roomKey
	r.DELETE("/:userKey/room_out/:roomKey", roomUserController.RoomOut()) // room/:userKey/room_out/:roomKey

	e.Logger.Fatal(e.Start(":8000"))
}
