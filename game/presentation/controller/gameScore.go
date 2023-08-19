package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type GameScoreController interface {
	ListGameScore() echo.HandlerFunc
	UpdateGameScore() echo.HandlerFunc
}

type gameScoreController struct {
	gameScoreService service.GameScoreService
}

func NewGameScoreController(
		gameScoreService service.GameScoreService,
	) GameScoreController {
    return &gameScoreController{
        gameScoreService: gameScoreService,
    }
}

// List
// @Summary     スコア更新
// @tags        GameScore
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGameScore}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      user/list_game_score [post]
func (gameScoreController *gameScoreController) ListGameScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		gameKey := c.Param("gameKey")
		userKey := c.Param("userKey")

		gameAndGameScoreResult, err := gameScoreController.gameScoreService.ListGameScore(gameKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game_score", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListGameScore(gameAndGameScoreResult)
		response := response.SuccessWith("list_game_score", 200, out)

		return c.JSON(200, response)
	}
}

// Update
// @Summary     スコア更新
// @tags        GameScore
// @Accept      json
// @Produce     json
// @Param       body body parameter.UpdateGameScore true "スコア更新"
// @Success     200  {object} response.Success{items=output.UpdateGameScore}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      user/update_game_score [post]
func (gameScoreController *gameScoreController) UpdateGameScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		gameScoreParam := &parameter.UpdateGameScore{}
		c.Bind(gameScoreParam)

		baseToken := c.Request().Header.Get("Authorization")
		userKey, _, _, err := CheckToken(baseToken)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("update_game_score", 500, out)

			return c.JSON(500, response)
		}

		gameScoreResult, err := gameScoreController.gameScoreService.UpdateGameScore(userKey, gameScoreParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("update_game_score", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToUpdateGameScore(gameScoreResult)
		response := response.SuccessWith("update_game_score", 200, out)

		return c.JSON(200, response)
	}
}
