package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type RoomController interface {
	ListRoom() echo.HandlerFunc
	SearchRoom() echo.HandlerFunc
	CreateRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
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
// @Success     200  {object} response.Success{items=output.ListRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/list_room [get]
func (roomController *roomController) ListRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")

		roomResult, err := roomController.roomService.ListRoom(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_room", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListRoom(roomResult)
		response := response.SuccessWith("list_room", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ルーム検索
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/search_room [get]
func (roomController *roomController) SearchRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		name := c.QueryParam("name")
		genre := c.QueryParam("genre")
		game := c.QueryParam("game")

		roomResult, err := roomController.roomService.SearchRoom(name, genre, game)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("search_room", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListRoom(roomResult)
		response := response.SuccessWith("search_room", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     ルーム作成
// @tags        Room
// @Accept      json
// @Produce     json
// @Param       body body parameter.CreateRoom true "ルーム作成"
// @Success     200  {object} response.Success{items=output.CreateRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /user/{user_key}/create_room  [post]
func (roomController *roomController) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		roomParam := &parameter.CreateRoom{}
		c.Bind(roomParam)

		roomResult, err := roomController.roomService.CreateRoom(roomParam, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_room", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateRoom(roomResult)
		response := response.SuccessWith("create_room", 200, out)

		return c.JSON(200, response)
	}
}

// Delete
// @Summary     ルーム削除
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.DeleteRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{user_key}/delete_room/{roomKey}  [delete]
func (roomController *roomController) DeleteRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		roomKey := c.Param("roomKey")

		err := roomController.roomService.DeleteRoom(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("delete_room", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToDeleteRoom()
		response := response.SuccessWith("delete_room", 200, out)
		
		return c.JSON(200, response)
	}
}
