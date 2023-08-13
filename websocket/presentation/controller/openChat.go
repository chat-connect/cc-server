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

type OpenChatController interface {
	SendOpenChat() echo.HandlerFunc
}

type openChatController struct {
	openChatService service.OpenChatService
	userService     service.UserService
}

func NewOpenChatController(
		openChatService service.OpenChatService,
		userService     service.UserService,
	) OpenChatController {
    return &openChatController{
		openChatService: openChatService,
		userService:     userService,
    }
}


// WebSocketでチャットを送信
func (openChatController *openChatController) SendOpenChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		openKey := "open"

		conn, err := roomChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil { return err }

		if rooms[openKey] == nil {
			rooms[openKey] = make(map[*websocket.Conn]bool)
		}

		rooms[openKey][conn] = true

		go func() {
			defer func() {
				conn.Close()
				delete(rooms[openKey], conn)
			}()

			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil { return }

				message := string(p)
				openChatParam := &parameter.CreateOpenChat{}
				err = json.Unmarshal([]byte(message), openChatParam)
				if err != nil { return }

				token := strings.ReplaceAll(openChatParam.Token, "Bearer ", "")
				userKey := openChatParam.UserKey
			
				user, err := openChatController.userService.FindByUserKey(userKey)
				if err != nil { return }
				if token != user.Token { return }

				openResult, err := openChatController.openChatService.CreateOpenChat(userKey, openChatParam)
				if err != nil { return }

				out := output.ToCreateOpenChat(openResult)
				response := response.SuccessWith("chat_create", 200, out)

				jsonResponse, err := json.Marshal(response)
				if err != nil { return }

				for client := range rooms[openKey] {
					err := client.WriteMessage(messageType, jsonResponse)
					if err != nil {
						delete(rooms[openKey], client)
					}
				}
			}
		}()

		return nil
	}
}
