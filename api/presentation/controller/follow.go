package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type FollowController interface {
	CountFollowingAndFollowers() echo.HandlerFunc
	ListFollowing() echo.HandlerFunc
	ListFollowers() echo.HandlerFunc
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
// @Summary     フォローしているユーザー一覧
// @tags        Follow
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.FollowingAndFollowers}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      follow/{userKey}/count_following_and_followers [get]
func (followController *followController) CountFollowingAndFollowers() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")

		followingCount, err := followController.followService.CountFollowing(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("count_following_and_followers", 500, out)

			return c.JSON(500, response)
		}

		followersCount, err := followController.followService.CountFollowers(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("count_following_and_followers", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCountFollowingAndFollowers(userKey, followingCount, followersCount)
		response := response.SuccessWith("count_following_and_followers", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     フォローしているユーザー一覧
// @tags        Follow
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListFollows}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      follow/{userKey}/list_following [post]
func (followController *followController) ListFollowing() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")

		followAndUserResults, err := followController.followService.ListFollowing(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_following", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListFollowing(userKey, followAndUserResults)
		response := response.SuccessWith("list_following", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     フォローしているユーザー一覧
// @tags        Follow
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListFollows}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      follow/{userKey}/list_followers [post]
func (followController *followController) ListFollowers() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")

		followAndUserResults, err := followController.followService.ListFollowers(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_followers", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListFollowers(userKey, followAndUserResults)
		response := response.SuccessWith("list_followers", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     フォロー作成
// @tags        Follow
// @Accept      json
// @Produce     json
// @Param       body body parameter.CreateFollow true "フォロー作成"
// @Success     200  {object} response.Success{items=output.CreateFollow}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      follow/{userKey}/create_follow [post]
func (followController *followController) CreateFollow() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		followParam := &parameter.CreateFollow{}
		c.Bind(followParam)

		// validation
		check, _ := followController.followService.FindByUserKeyAndFollowingUserKey(userKey, followParam.FollowingUserKey)
		if check != nil && check.UserKey == userKey && check.FollowingUserKey ==followParam.FollowingUserKey {
			out := output.NewError(fmt.Errorf("already followed"))
			response := response.ErrorWith("chat_follow", 404, out)

			return c.JSON(404, response)
		}

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
