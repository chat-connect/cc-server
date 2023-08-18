package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type LinkGameController interface {
	CreateLinkGame() echo.HandlerFunc
}

type linkGameController struct {
	linkGameService service.LinkGameService
}

func NewLinkGameController(linkGameService service.LinkGameService) LinkGameController {
    return &linkGameController{
        linkGameService: linkGameService,
    }
}

// Create
// @Summary     連携ゲーム作成
// @tags        LinkGame
// @Accept      json
// @Produce     json
// @Param       body body parameter.CreateLinkGame true "連携ゲーム作成"
// @Success     200  {object} response.Success{items=output.CreateLinkGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /ink_game/{admin_user_key}/create_ink_game  [post]
func (linkGameController *linkGameController) CreateLinkGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		adminUserKey := c.Param("adminUserKey")
		linkGameParam := &parameter.CreateLinkGame{}
		c.Bind(linkGameParam )

		linkGameResult, err := linkGameController.linkGameService.CreateLinkGame(adminUserKey, linkGameParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_link_game", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateLinkGame(linkGameResult)
		response := response.SuccessWith("create_link_game", 200, out)

		return c.JSON(200, response)
	}
}
