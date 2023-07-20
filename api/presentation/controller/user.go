package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
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
// @Success     200  {object} output.UserRegister
// @Failure     500  {array}  output.Error
// @Router      /auth/user_register [post]
func (userController *userController) UserRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)

		// 登録済のメールアドレスを検索
		_, err := userController.userService.FindByEmail(userModel.Email)
		if err == nil {
			return c.JSON(400, output.ToEmailValidation())
		}

		userResult, err := userController.userService.UserRegister(userModel)
		if err != nil {
			return c.JSON(500, output.NewError(err))
		}

		return c.JSON(200, output.ToUserRegister(userResult))
	}
}

// Login
// @Summary     ユーザーログイン
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.UserLogin true "ユーザーログイン"
// @Success     200  {object} output.UserLogin
// @Failure     500  {array}  output.Error
// @Router      /auth/user_login [post]
func (userController *userController) UserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)

		user, err := userController.userService.UserLogin(userModel)
		if err != nil {
			return c.JSON(500, output.NewError(err))
		}
	
		return c.JSON(200, output.ToUserLogin(user))
	}
}

// Check
// @Summary     ユーザー取得
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Success     200  {object} output.UserCheck
// @Failure     500  {array}  output.Error
// @Router      /user/{user_key}/user_check [get]
func (userController *userController) UserCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		userKey, username, email, err := userController.userService.UserCheck(baseToken)
		if err != nil {
			return c.JSON(500, output.NewError(err))
		}
	
		return c.JSON(200, output.ToUserCheck(userKey, username, email))
	}
}

// Logout
// @Summary     ログアウト
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path string true "ユーザーキー"
// @Success     200  {object} output.UserLogout
// @Failure     500  {array}  output.Error
// @Router      /user/{user_key}/user_logout [put]
func (userController *userController) UserLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		userModel := &model.User{}
		c.Bind(userModel)
		
		userKey := c.Param("userKey")
		userModel.UserKey =  userKey 
		userModel.Status = "offline"
		userModel.Token = "nil"
		_, err := userController.userService.UserLogout(userModel)
		if err != nil {
			return c.JSON(500, output.NewError(err))
		}

		return c.JSON(200, output.ToUserLogout())
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
// @Success     200  {object} output.UserDelete
// @Failure     500  {array}  output.Error
// @Router      /user/{user_key}/user_delete [delete]
func (userController *userController) UserDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userKey := c.Param("userKey")
	
		err := userController.userService.UserDelete(userKey)
		if err != nil {
			return c.JSON(500, output.NewError(err))
		}

		return c.JSON(200, output.ToUserDelete())
	}
}
