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

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// realtime: 同期関連
	chat := e.Group("/realtime")
	chat.GET("/:channelKey/send_chat", chatController.SendChat()) // realtime/:channelKey/chat_create

	e.Logger.Fatal(e.Start(":8000"))
}