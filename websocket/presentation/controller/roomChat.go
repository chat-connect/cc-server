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

type RoomChatController interface {
	SendRoomChat() echo.HandlerFunc
}

type roomChatController struct {
	roomChatService service.RoomChatService
	userService     service.UserService
}

func NewRoomChatController(
	roomChatService service.RoomChatService,
	userService     service.UserService,
) RoomChatController {
	return &roomChatController{
		roomChatService: roomChatService,
		userService:     userService,
	}
}

// WebSocketでチャットを送信
func (roomChatController *roomChatController) SendRoomChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomKey := c.Param("roomKey")

		conn, err := roomChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil { return err }

		if rooms[roomKey] == nil {
			rooms[roomKey] = make(map[*websocket.Conn]bool)
		}

		rooms[roomKey][conn] = true

		go func() {
			defer func() {
				conn.Close()
				delete(rooms[roomKey], conn)
			}()

			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil { return }

				message := string(p)
				roomChatParam := &parameter.CreateRoomChat{}
				err = json.Unmarshal([]byte(message), roomChatParam)
				if err != nil { return }

				token := strings.ReplaceAll(roomChatParam.Token, "Bearer ", "")
				userKey := roomChatParam.UserKey
			
				user, err := roomChatController.userService.FindByUserKey(userKey)
				if err != nil { return }
				if token != user.Token { return }

				roomResult, err := roomChatController.roomChatService.CreateRoomChat(roomKey, userKey, roomChatParam)
				if err != nil { return }

				out := output.ToCreateRoomChat(roomResult)
				response := response.SuccessWith("chat_create", 200, out)

				jsonResponse, err := json.Marshal(response)
				if err != nil { return }

				for client := range rooms[roomKey] {
					err := client.WriteMessage(messageType, jsonResponse)
					if err != nil {
						delete(rooms[roomKey], client)
					}
				}
			}
		}()

		return nil
	}
}
