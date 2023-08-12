package controller

import (
	"strings"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/websocket"
	
	"github.com/game-connect/gc-server/websocket/service"
	"github.com/game-connect/gc-server/websocket/presentation/output"
	"github.com/game-connect/gc-server/websocket/presentation/response"
	"github.com/game-connect/gc-server/websocket/presentation/parameter"
)

type ChatController interface {
	SendChat() echo.HandlerFunc
}

type chatController struct {
	chatService service.ChatService
	userService service.UserService
}

func NewChatController(
	chatService service.ChatService,
	userService service.UserService,
) ChatController {
	return &chatController{
		chatService: chatService,
		userService: userService,
	}
}

var (
	rooms = make(map[string]map[*websocket.Conn]bool)
	roomChatUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

// WebSocketでチャットを送信
func (chatController *chatController) SendChat() echo.HandlerFunc {
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
				chatParam := &parameter.CreateChat{}
				err = json.Unmarshal([]byte(message), chatParam)
				if err != nil { return }

				token := strings.ReplaceAll(chatParam.Token, "Bearer ", "")
				userKey := chatParam.UserKey
			
				user, err := chatController.userService.FindByUserKey(userKey)
				if err != nil { return }
				if token != user.Token { return }

				chatResult, err := chatController.chatService.CreateChat(channelKey, userKey, chatParam)
				if err != nil { return }

				out := output.ToCreateChat(chatResult)
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
