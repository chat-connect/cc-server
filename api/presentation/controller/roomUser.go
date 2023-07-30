package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
)

type RoomUserController interface {
	RoomJoin() echo.HandlerFunc
}

type roomUserController struct {
	roomUserService service.RoomUserService
}

func NewRoomUserController(
		roomUserService service.RoomUserService,
	) RoomUserController {
    return &roomUserController{
		roomUserService: roomUserService,
    }
}

// Join
// @Summary     ルーム参加
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{Items=output.RoomJoin}
// @Failure     500  {array}  output.Error
// @Router      /user/{userKey}/room_join/{roomKey}  [post]
func (roomUserController *roomUserController) RoomJoin() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")

		roomResult, err := roomUserController.roomUserService.RoomJoin(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_join", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRoomJoin(roomResult)
		response := response.SuccessWith("room_join", 200, out)

		return c.JSON(200, response)
	}
}
