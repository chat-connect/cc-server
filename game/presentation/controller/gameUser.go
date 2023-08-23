package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
)

type GameUserController interface {
	ListGameUser() echo.HandlerFunc
}

type gameUserController struct {
	gameUserService service.GameUserService
}

func NewGameUserController(gameUserService service.GameUserService) GameUserController {
    return &gameUserController{
        gameUserService: gameUserService,
    }
}

// List
// @Summary     連携ゲーム一覧取得
// @tags        GameUser
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGameUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /game/{userKey}/list_game_user [get]
func (gameUserController *gameUserController) ListGameUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
		gameUserResult, err := gameUserController.gameUserService.ListGameUser(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListGameUser(gameUserResult)
		response := response.SuccessWith("list_game_user", 200, out)

		return c.JSON(200, response)
	}
}
