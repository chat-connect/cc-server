package controller

import (
	"strings"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/websocket"
	
	"github.com/game-connect/gc-server/websocket/service"
	"github.com/game-connect/gc-server/websocket/presentation/output"
	"github.com/game-connect/gc-server/websocket/presentation/response"
	"github.com/game-connect/gc-server/websocket/presentation/parameter"
)

type ChannelChatController interface {
	SendChannelChat() echo.HandlerFunc
}

type channelChatController struct {
	channelChatService service.ChannelChatService
	userService service.UserService
}

func NewChannelChatController(
	channelChatService service.ChannelChatService,
	userService        service.UserService,
) ChannelChatController {
	return &channelChatController{
		channelChatService: channelChatService,
		userService:        userService,
	}
}

// WebSocketでチャットを送信
func (channelChatController *channelChatController) SendChannelChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		channelKey := c.Param("channelKey")

		conn, err := roomChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil { return err }

		if rooms[channelKey] == nil {
			rooms[channelKey] = make(map[*websocket.Conn]bool)
		}

		rooms[channelKey][conn] = true

		go func() {
			defer func() {
				conn.Close()
				delete(rooms[channelKey], conn)
			}()

			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil { return }

				message := string(p)
				channelChatParam := &parameter.CreateChannelChat{}
				err = json.Unmarshal([]byte(message), channelChatParam)
				if err != nil { return }

				token := strings.ReplaceAll(channelChatParam.Token, "Bearer ", "")
				userKey := channelChatParam.UserKey
			
				user, err := channelChatController.userService.FindByUserKey(userKey)
				if err != nil { return }
				if token != user.Token { return }

				chatResult, err := channelChatController.channelChatService.CreateChannelChat(channelKey, userKey, channelChatParam)
				if err != nil { return }

				out := output.ToCreateChannelChat(chatResult)
				response := response.SuccessWith("chat_create", 200, out)

				jsonResponse, err := json.Marshal(response)
				if err != nil { return }

				for client := range rooms[channelKey] {
					err := client.WriteMessage(messageType, jsonResponse)
					if err != nil {
						delete(rooms[channelKey], client)
					}
				}
			}
		}()

		return nil
	}
}
