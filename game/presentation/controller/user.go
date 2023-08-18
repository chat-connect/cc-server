package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type UserController interface {
	RegisterUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	CheckUser() echo.HandlerFunc
	LogoutUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
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
// @Summary     ユーザー登録
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.RegisterUser true "ユーザー登録"
// @Success     200  {object} response.Success{items=output.RegisterUser}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_register [post]
func (userController *userController) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userParam := &parameter.RegisterUser{}
		c.Bind(userParam)

		// 登録済のメールアドレスを検索
		check, _ := userController.userService.FindByEmail(userParam.Email)
		if check.Email == userParam.Email {
			response := response.ErrorWith("user_register", 400, "email already exists")

			return c.JSON(400, response)
		}

		userResult, err := userController.userService.RegisterUser(userParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_register", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRegisterUser(userResult)
		response := response.SuccessWith("user_register", 200, out)

		return c.JSON(200, response)
	}
}

// Login
// @Summary     ユーザーログイン
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.LoginUser true "ユーザーログイン"
// @Success     200  {object} response.Success{items=output.LoginUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /auth/user_login [post]
func (userController *userController) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)

		userResult, err := userController.userService.LoginUser(userModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_login", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToLoginUser(userResult)
		response := response.SuccessWith("user_login", 200, out)
		
		return c.JSON(200, response)
	}
}

// Check
// @Summary     ユーザー取得
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path string true "user_key" maxlength(12)
// @Success     200  {object} response.Success{items=output.CheckUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /auth/user_check/{user_key} [get]
func (userController *userController) CheckUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		userKey, name, email, err := userController.userService.CheckUser(baseToken)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_check", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToCheckUser(userKey, name, email)
		response := response.SuccessWith("user_check", 200, out)
		
		return c.JSON(200, response)
	}
}

// Logout
// @Summary     ログアウト
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path string true "user_key" maxlength(12)
// @Success     200  {object} response.Success{items=output.LogoutUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /auth/user_logout/{user_key} [put]
func (userController *userController) LogoutUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)
		
		userKey := c.Param("userKey")
		userModel.UserKey =  userKey 
		userModel.Token = "nil"
		_, err := userController.userService.LogoutUser(userModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_logout", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToLogoutUser()
		response := response.SuccessWith("user_logout", 200, out)
		
		return c.JSON(200, response)
	}
}

// Delete
// @Summary     ユーザー削除
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path string true "ユーザーキー"
// @Success     200  {object} response.Success{items=output.DeleteUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /auth/user_delete/{user_key} [delete]
func (userController *userController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
		err := userController.userService.DeleteUser(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_delete", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToDeleteUser()
		response := response.SuccessWith("user_delete", 200, out)
		
		return c.JSON(200, response)
	}
}
