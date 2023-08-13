package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/game-connect/gc-server/log"
	"github.com/game-connect/gc-server/websocket/di"
)

func Init() {
	// di: wire ./api/di/wire.go
	chatController := di.InitializeChatController()
	roomChatController := di.InitializeRoomChatController()
	openChatController := di.InitializeOpenChatController()
	channelChatController := di.InitializeChannelChatController()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// realtime: 同期関連
	chat := e.Group("/realtime")
	chat.GET("/:channelKey/send_chat", chatController.SendChat()) // realtime/:channelKey/send_chat
	chat.GET("/send_open_chat", openChatController.SendOpenChat()) // realtime/send_open_chat
	chat.GET("/:roomKey/send_room_chat", roomChatController.SendRoomChat()) // realtime/:roomKey/send_room_chat
	chat.GET("/:channelKey/send_channel_chat", channelChatController.SendChannelChat()) // realtime/:channelKey/send_channel_chat

	e.Logger.Fatal(e.Start(":8000"))
}
