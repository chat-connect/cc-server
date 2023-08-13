package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type RoomChatController interface {
	ListRoomChat() echo.HandlerFunc
	CreateRoomChat() echo.HandlerFunc
}

type roomChatController struct {
	roomChatService service.RoomChatService
}

func NewRoomChatController(
		roomChatService service.RoomChatService,
	) RoomChatController {
    return &roomChatController{
		roomChatService: roomChatService,
    }
}

// List
// @Summary     チャット一覧取得
// @tags        RoomChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListRoomChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/list_croom_chat/{roomKey} [get]
func (roomChatController *roomChatController) ListRoomChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")

		roomChatResult, err := roomChatController.roomChatService.ListRoomChat(roomKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_room_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListRoomChat(roomKey, roomChatResult)
		response := response.SuccessWith("list_room_chat", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     チャット作成
// @tags        RoomChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateRoomChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/create_room_chat/{roomKey} [post]
func (roomChatController *roomChatController) CreateRoomChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")
		roomChatParam := &parameter.CreateRoomChat{}
		c.Bind(roomChatParam)

		roomChatResult, err := roomChatController.roomChatService.CreateRoomChat(roomKey, userKey, roomChatParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_room_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateRoomChat(roomChatResult)
		response := response.SuccessWith("create_room_chat", 200, out)

		return c.JSON(200, response)
	}
}
