package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/game-connect/gc-server/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-connect/gc-server/log"
	"github.com/game-connect/gc-server/game/di"
)

func Init() {
	// di: wire ./api/di/wire.go
	adminUserController := di.InitializeAdminUserController()
	adminUserMiddleware := di.InitializeAdminUserMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth: 認証API
	auth := e.Group("/admin")
	auth.POST("/register_admin_user", adminUserController.RegisterAdminUser()) // admin/register_admin_use
	auth.POST("/login_admin_user", adminUserController.LoginAdminUser()) // auth/user_login

	// auth: 認証済ユーザーのみアクセス可能
	auth.Use(adminUserMiddleware.AdminUserMiddleware)
	auth.GET("/check_admin_user/:adminUserKey", adminUserController.CheckAdminUser()) // auth/user_check/:userKey
	auth.PUT("/logout_admin_user/:adminUserKey", adminUserController.LogoutAdminUser()) // auth/user_logout/:userKey
	auth.DELETE("/delete_admin_user/:adminUserKey", adminUserController.DeleteAdminUser()) // auth/user_delete/:userKey

	e.Logger.Fatal(e.Start(":8000"))
}
