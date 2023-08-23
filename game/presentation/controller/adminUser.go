package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type AdminUserController interface {
	RegisterAdminUser() echo.HandlerFunc
	LoginAdminUser() echo.HandlerFunc
	CheckAdminUser() echo.HandlerFunc
	LogoutAdminUser() echo.HandlerFunc
	DeleteAdminUser() echo.HandlerFunc
}

type adminUserController struct {
	adminUserService service.AdminUserService
}

func NewAdminUserController(adminUserService service.AdminUserService) AdminUserController {
    return &adminUserController{
        adminUserService: adminUserService,
    }
}

// Register
// @Summary     企業ユーザー登録
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.RegisterAdminUser true "企業ユーザー登録"
// @Success     200  {object} response.Success{items=output.RegisterAdminUser}
// @Failure     500  {array}  output.Error
// @Router      /admin/user_admin_register [post]
func (adminUserController *adminUserController) RegisterAdminUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminUserParam := &parameter.RegisterAdminUser{}
		c.Bind(adminUserParam)

		// 登録済のメールアドレスを検索
		check, _ := adminUserController.adminUserService.FindByEmail(adminUserParam.Email)
		if check.Email == adminUserParam.Email {
			response := response.ErrorWith("register_admin_user", 400, "email already exists")

			return c.JSON(400, response)
		}

		adminUserResult, err := adminUserController.adminUserService.RegisterAdminUser(adminUserParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("register_admin_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToRegisterAdminUser(adminUserResult)
		response := response.SuccessWith("register_admin_user", 200, out)

		return c.JSON(200, response)
	}
}

// Login
// @Summary     企業ユーザーログイン
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body parameter.LoginUser true "企業ユーザーログイン"
// @Success     200  {object} response.Success{items=output.LoginAdminUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /admin/login_admin_user [post]
func (adminUserController *adminUserController) LoginAdminUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminUserParam := &parameter.LoginAdminUser{}
		c.Bind(adminUserParam)

		adminUserResult, err := adminUserController.adminUserService.LoginAdminUser(adminUserParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("user_login", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToLoginAdminUser(adminUserResult)
		response := response.SuccessWith("user_login", 200, out)
		
		return c.JSON(200, response)
	}
}

// Check
// @Summary     企業ユーザー取得
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path string true "user_key" maxlength(12)
// @Success     200  {object} response.Success{items=output.CheckAdminUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /admin/check_admin_user/{admin_user_key} [get]
func (adminUserController *adminUserController) CheckAdminUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		adminUserKey, name, email, err := adminUserController.adminUserService.CheckAdminUser(baseToken)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("check_admin_user", 500, out)

			return c.JSON(500, response)
		}
	
		out := output.ToCheckAdminUser(adminUserKey, name, email)
		response := response.SuccessWith("check_admin_user", 200, out)
		
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
// @Success     200  {object} response.Success{items=output.LogoutAdminUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /auth/user_logout/{user_key} [put]
func (adminUserController *adminUserController) LogoutAdminUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminUserModel := &model.AdminUser{}
		c.Bind(adminUserModel)
		
		adminUserKey := c.Param("adminUserKey")
		adminUserModel.AdminUserKey = adminUserKey
		adminUserModel.Token = "nil"
		_, err := adminUserController.adminUserService.LogoutAdminUser(adminUserModel)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("logout_admin_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToLogoutAdminUser()
		response := response.SuccessWith("logout_admin_user", 200, out)
		
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
// @Success     200  {object} response.Success{items=output.DeleteAdminUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /admin/delete_admin_user/{admin_user_key} [delete]
func (adminUserController *adminUserController) DeleteAdminUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminUserKey := c.Param("adminUserKey")
		err := adminUserController.adminUserService.DeleteAdminUser(adminUserKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("delete_admin_user", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToDeleteAdminUser()
		response := response.SuccessWith("delete_admin_user", 200, out)
		
		return c.JSON(200, response)
	}
}
