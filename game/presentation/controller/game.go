package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type GameController interface {
	CreateGame() echo.HandlerFunc
}

type gameController struct {
	gameService service.GameService
}

func NewGameController(gameService service.GameService) GameController {
    return &gameController{
        gameService: gameService,
    }
}

// Create
// @Summary     連携ゲーム作成
// @tags        LinkGame
// @Accept      json
// @Produce     json
// @Param       body body parameter.CreateGame true "連携ゲーム作成"
// @Success     200  {object} response.Success{items=output.CreateGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /game/{admin_user_key}/create_game  [post]
func (gameController *gameController) CreateGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		adminUserKey := c.Param("adminUserKey")
		gameParam := &parameter.CreateGame{}
		c.Bind(gameParam)

		linkGameResult, err := gameController.gameService.CreateGame(adminUserKey, gameParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_game", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateGame(linkGameResult)
		response := response.SuccessWith("create_game", 200, out)

		return c.JSON(200, response)
	}
}
