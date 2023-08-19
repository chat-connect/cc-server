package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type UserController interface {
	LoginUser() echo.HandlerFunc
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
    return &userController{
        userService: userService,
    }
}

// Login
// @Summary     ユーザーログイン
// @tags        User
// @Accept      json
// @Produce     json
// @Param       body body parameter.LoginUser true "ユーザーログイン"
// @Success     200  {object} response.Success{items=output.LoginUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /user/login_user [post]
func (userController *userController) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userParam := &parameter.LoginUser{}
		c.Bind(userParam)

		userResult, err := userController.userService.LoginUser(userParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("login_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToLoginUser(userResult)
		response := response.SuccessWith("login_user", 200, out)
		
		return c.JSON(200, response)
	}
}
