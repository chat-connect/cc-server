package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
)

type RoomController interface {
	RoomCreate() echo.HandlerFunc
}

type roomController struct {
	roomService service.RoomService
}

func NewRoomController(roomService service.RoomService) RoomController {
    return &roomController{
        roomService: roomService,
    }
}

// Create
// @Summary     ルーム作成
// @tags        Room
// @Accept      json
// @Produce     json
// @Param       body body parameter.RoomCreate true "ルーム作成"
// @Success     200  {object} response.Success{items=output.RoomCreate}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /user/{user_key}/room_create  [post]
func (roomController *roomController) RoomCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		roomParam := &parameter.RoomCreate{}
		c.Bind(roomParam)

		roomResult, err := roomController.roomService.RoomCreate(roomParam, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_create", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRoomCreate(roomResult)
		response := response.SuccessWith("room_create", 200, out)

		return c.JSON(200, response)
	}
}
