package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
)

type UserController interface {
	UserRegister() echo.HandlerFunc
	UserLogin() echo.HandlerFunc
	UserCheck() echo.HandlerFunc
	UserLogout() echo.HandlerFunc
	UserDelete() echo.HandlerFunc
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
// @Param       body body parameter.UserRegister true "ユーザー登録"
// @Success     200  {object} response.Success{Items=output.UserRegister}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_register [post]
func (userController *userController) UserRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)

		// 登録済のメールアドレスを検索
		check, _ := userController.userService.FindByEmail(userModel.Email)
		if check.Email == userModel.Email {
			response := response.ErrorWith("user_register", 400, "email already exists")

			return c.JSON(400, response)
		}

		userResult, err := userController.userService.UserRegister(userModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_register", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToUserRegister(userResult)
		response := response.SuccessWith("user_register", 200, out)

		return c.JSON(200, response)
	}
}

// Login
// @Summary     ユーザーログイン
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.UserLogin true "ユーザーログイン"
// @Success     200  {object} response.Success{Items=output.UserLogin}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_login [post]
func (userController *userController) UserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)

		userResult, err := userController.userService.UserLogin(userModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_login", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToUserLogin(userResult)
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
// @Success     200  {object} response.Success{Items=output.UserCheck}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_check/{user_key} [get]
func (userController *userController) UserCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		userKey, name, email, err := userController.userService.UserCheck(baseToken)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_check", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToUserCheck(userKey, name, email)
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
// @Success     200  {object} response.Success{Items=output.UserLogout}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_logout/{user_key} [put]
func (userController *userController) UserLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)
		
		userKey := c.Param("userKey")
		userModel.UserKey =  userKey 
		userModel.Status = "logout"
		userModel.Token = "nil"
		_, err := userController.userService.UserLogout(userModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_logout", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToUserLogout()
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
// @Success     200  {object} response.Success{Items=output.UserDelete}
// @Failure     500  {array}  output.Error
// @Router      /auth/user_delete/{user_key} [delete]
func (userController *userController) UserDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
		err := userController.userService.UserDelete(userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_delete", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToUserDelete()
		response := response.SuccessWith("user_delete", 200, out)
		
		return c.JSON(200, response)
	}
}
