package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
)

type RoomUserController interface {
	JoinRoom() echo.HandlerFunc
	OutRoom() echo.HandlerFunc
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
// @Success     200  {object} response.Success{items=output.JoinRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/room_join/{roomKey}  [post]
func (roomUserController *roomUserController) JoinRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")

		roomResult, err := roomUserController.roomUserService.JoinRoom(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_join", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToJoinRoom(roomResult)
		response := response.SuccessWith("room_join", 200, out)

		return c.JSON(200, response)
	}
}

// Out
// @Summary     ルーム退出
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.OutRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/room_out/{roomKey}  [delete]
func (roomUserController *roomUserController) OutRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")

		err := roomUserController.roomUserService.OutRoom(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_out", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToOutRoom()
		response := response.SuccessWith("room_out", 200, out)

		return c.JSON(200, response)
	}
}
