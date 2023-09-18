package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

type UserController interface {
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	SearchUser() echo.HandlerFunc
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
    return &userController{
        userService: userService,
    }
}

// Get
// @Summary     ユーザー情報
// @tags        User
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.GetUser}
// @Failure     500  {array}  output.Error
// @Router      /user/{userKey}/get_user [get]
func (userController *userController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")

		userResult, err := userController.userService.FindByUserKey(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("get_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToGetUser(userResult)
		response := response.SuccessWith("get_user", 200, out)

		return c.JSON(200, response)
	}
}

// @Summary     ユーザー更新
// @tags        User
// @Accept      json
// @Produce     json
// @Param       body body parameter.UpdateUser true "ユーザー更新"
// @Success     200  {object} response.Success{items=output.GetUser}
// @Failure     500  {array}  output.Error
// @Router      /user/{userKey}/update_user [get]
func (userController *userController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
		userParam := &parameter.UpdateUser{}
		c.Bind(userParam)

		userResult, err := userController.userService.UpdateUser(userKey, userParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("update_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToUpdateUser(userResult)
		response := response.SuccessWith("update_user", 200, out)

		return c.JSON(200, response)
	}
}

// Serach
// @Summary     ユーザー検索
// @tags        User
// @Accept      json
// @Produce     json
// @Param       body body parameter.SearchUser true "ユーザー検索"
// @Success     200  {object} response.Success{items=output.SearchUser}
// @Failure     500  {array}  output.Error
// @Router      /user/{userKey}/serach_user [get]
func (userController *userController) SearchUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
		name := c.QueryParam("name")

		userResults, err := userController.userService.SearchUser(userKey, name)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("search_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToSearchUser(userResults)
		response := response.SuccessWith("search_user", 200, out)

		return c.JSON(200, response)
	}
}
