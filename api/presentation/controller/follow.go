package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type FollowController interface {
	CreateFollow() echo.HandlerFunc
}

type followController struct {
	followService service.FollowService
}

func NewFollowController(
		followService service.FollowService,
	) FollowController {
    return &followController{
		followService: followService,
    }
}

// Create
// @Summary     フォロー作成
// @tags        Follow
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateFollow}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      direct_mail/{userKey}/create_direct_mail [post]
func (followController *followController) CreateFollow() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		followParam := &parameter.CreateFollow{}
		c.Bind(followParam)

		followResult, err := followController.followService.CreateFollow(userKey, followParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("chat_follow", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateFollow(followResult)
		response := response.SuccessWith("chat_follow", 200, out)

		return c.JSON(200, response)
	}
}
