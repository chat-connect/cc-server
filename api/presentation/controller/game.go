package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

type GameController interface {
	ListGame() echo.HandlerFunc
}

type gameController struct {
	gameService service.GameService
}

func NewGameController(
		gameService service.GameService,
	) GameController {
    return &gameController{
		gameService: gameService,
    }
}

// List
// @Summary     ジャンル一覧取得
// @tags        Game
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_game [get]
func (gameController *gameController) ListGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		gameResult, err := gameController.gameService.ListGame()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGame(gameResult)
		response := response.SuccessWith("list_game", 200, out)

		return c.JSON(200, response)
	}
}
