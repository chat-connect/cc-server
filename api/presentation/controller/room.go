package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
)

type RoomController interface {
	RoomList() echo.HandlerFunc
	RoomCreate() echo.HandlerFunc
	RoomDelete() echo.HandlerFunc
}

type roomController struct {
	roomService service.RoomService
}

func NewRoomController(roomService service.RoomService) RoomController {
    return &roomController{
        roomService: roomService,
    }
}

// List
// @Summary     ルーム一覧取得
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.RoomList}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/room_list [get]
func (roomController *roomController) RoomList() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")

		roomResult, err := roomController.roomService.RoomList(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_list", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRoomList(roomResult)
		response := response.SuccessWith("room_list", 200, out)

		return c.JSON(200, response)
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

// Delete
// @Summary     ルーム削除
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.RoomDelete}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{user_key}/room_delete/{roomKey}  [delete]
func (roomController *roomController) RoomDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		roomKey := c.Param("roomKey")

		err := roomController.roomService.RoomDelete(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_delete", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRoomDelete()
		response := response.SuccessWith("room_delete", 200, out)
		
		return c.JSON(200, response)
	}
}
