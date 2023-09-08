package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

type UserController interface {
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

// Register
// @Summary     ユーザー検索
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.SearchUser true "ユーザー検索"
// @Success     200  {object} response.Success{items=output.SearchUser}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_register [post]
func (userController *userController) SearchUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		userResults, err := userController.userService.SearchUser(name)
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
