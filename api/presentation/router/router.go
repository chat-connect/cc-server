package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/game-connect/gc-server/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-connect/gc-server/log"
	"github.com/game-connect/gc-server/api/di"
)

func Init() {
	// di: wire ./api/di/wire.go
	userController := di.InitializeUserController()
	roomController := di.InitializeRoomController()
	roomUserController := di.InitializeRoomUserController()
	channelController := di.InitializeChannelController()
	chatController := di.InitializeChatController()
	openChatController := di.InitializeOpenChatController()
	roomChatController := di.InitializeRoomChatController()
	channelChatController := di.InitializeChannelChatController()

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
	
	// open_chat: オープンチャット関連
	openChat := e.Group("/open_chat")
	openChat.Use(userMiddleware.UserMiddleware)
	openChat.GET("/:userKey/list_open_chat", openChatController.ListOpenChat()) // open_chat/:userKey/list_open_chat
	openChat.POST("/:userKey/create_open_chat", openChatController.CreateOpenChat()) // open_chat/:userKey/create_open_chat/:channelKey

	// room_chat: ルームチャット関連
	roomChat := e.Group("/room_chat")
	roomChat.Use(userMiddleware.UserMiddleware)
	roomChat.GET("/:userKey/list_room_chat/:roomKey", roomChatController.ListRoomChat()) // room_chat/:userKey/list_room_chat/:roomKey
	roomChat.POST("/:userKey/create_room_chat/:roomKey", roomChatController.CreateRoomChat()) // room_chat/:userKey/create_room_chat/:roomKey

	// channel_chat: チャンネルチャット関連
	channelChat := e.Group("/channel_chat")
	channelChat.Use(userMiddleware.UserMiddleware)
	channelChat.GET("/:userKey/list_channel_chat/:channelKey", channelChatController.ListChannelChat()) // channel_chat/:userKey/list_channel_chat/:channelKey
	channelChat.POST("/:userKey/create_channel_chat/:channelKey", channelChatController.CreateChannelChat()) // channel_chat/:userKey/create_channel_chat/:channelKey

	e.Logger.Fatal(e.Start(":8000"))
}
