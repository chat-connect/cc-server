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
	channelController := di.InitializeChannelController()
	chatController := di.InitializeChatController()

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

	// room: 部屋関連
	room := e.Group("/room")
	room.Use(userMiddleware.UserMiddleware)
	room.GET("/:userKey/list_room", roomController.ListRoom()) // room/:userKey/room_list
	room.POST("/:userKey/create_room", roomController.CreateRoom()) // room/:userKey/room_create
	room.DELETE("/:userKey/delete_room/:roomKey", roomController.DeleteRoom()) // room/:userKey/room_delete/:roomKey

	room.POST("/:userKey/join_room/:roomKey", roomUserController.JoinRoom()) // room/:userKey/room_join/:roomKey
	room.DELETE("/:userKey/out_room/:roomKey", roomUserController.OutRoom()) // room/:userKey/room_out/:roomKey

	// channel: チャンネル関連
	channel := e.Group("/channel")
	channel.Use(userMiddleware.UserMiddleware)
	channel.GET("/:userKey/list_channel/:roomKey", channelController.ListChannel()) // channel/:userKey/channel_list/:roomKey
	channel.POST("/:userKey/create_channel/:roomKey", channelController.CreateChannel()) // channel/:userKey/channel_create/:roomKey
	channel.DELETE("/:userKey/delete_channel/:channelKey", channelController.DeleteChannel()) // channel/:userKey/channel_delete/:channelKey

	// chat: チャット関連
	chat := e.Group("/chat")
	chat.Use(userMiddleware.UserMiddleware)
	chat.GET("/:userKey/list_chat/:channelKey", chatController.ListChat()) // chat/:userKey/chat_list/:channelKey
	chat.POST("/:userKey/create_chat/:channelKey", chatController.CreateChat()) // chat/:userKey/chat_create/:channelKey

	e.Logger.Fatal(e.Start(":8000"))
}
