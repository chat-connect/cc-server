package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

type GenreController interface {
	ListGenre() echo.HandlerFunc
	ListGame() echo.HandlerFunc
	ListGenreAndGame() echo.HandlerFunc
}

type genreController struct {
	genreService service.GenreService
	gameService service.GameService
}

func NewGenreController(
		genreService service.GenreService,
		gameService service.GameService,
	) GenreController {
    return &genreController{
		genreService: genreService,
		gameService:  gameService,
    }
}

// List
// @Summary     ジャンル一覧取得
// @tags        Genre
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGenre}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_genre [get]
func (genreController *genreController) ListGenre() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreResult, err := genreController.genreService.ListGenre()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_genre", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListGenre(genreResult)
		response := response.SuccessWith("list_genre", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ゲーム一覧取得
// @tags        Game
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_game [get]
func (genreController *genreController) ListGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		gameResult, err := genreController.gameService.ListGame()
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

// List
// @Summary     ジャンル＆ゲーム一覧取得
// @tags        Genre
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGenreAndGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_genre [get]
func (genreController *genreController) ListGenreAndGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreResult, err := genreController.genreService.ListGenre()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_genre_and_game", 500, out)

			return c.JSON(500, response)
		}

		gameResult, err := genreController.gameService.ListGame()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game_and_game", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGenreAndGame(genreResult, gameResult)
		response := response.SuccessWith("list_genre_and_game", 200, out)

		return c.JSON(200, response)
	}
}
