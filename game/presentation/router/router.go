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
	// di: wire ./game/di/wire.go
	adminUserController := di.InitializeAdminUserController()
	userController := di.InitializeUserController()
	linkGameController := di.InitializeLinkGameController()

	adminUserMiddleware := di.InitializeAdminUserMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// admin: 認証API
	admin := e.Group("/admin")
	admin.POST("/register_admin_user", adminUserController.RegisterAdminUser()) // admin/register_admin_use
	admin.POST("/login_admin_user", adminUserController.LoginAdminUser()) // admin/register_admin_user

	// admin: 認証済ユーザーのみアクセス可能
	admin.Use(adminUserMiddleware.AdminUserMiddleware)
	admin.GET("/check_admin_user/:adminUserKey", adminUserController.CheckAdminUser()) // admin/register_admin_check/:adminUserKey
	admin.PUT("/logout_admin_user/:adminUserKey", adminUserController.LogoutAdminUser()) // admin/register_admin_logout/:adminUserKey
	admin.DELETE("/delete_admin_user/:adminUserKey", adminUserController.DeleteAdminUser()) // admin/register_admin_delete/:adminUserKey

	// admin: ゲームを登録
	linkGame := e.Group("/link_game")
	linkGame.Use(adminUserMiddleware.AdminUserMiddleware)
	linkGame.POST("/:adminUserKey/create_link_game", linkGameController.CreateLinkGame()) // linkGame/:adminUserKey/create_link_game

	// user: 認証API 
	user := e.Group("/user")
	user.POST("/login_user", userController.LoginUser()) // user/login_user

	e.Logger.Fatal(e.Start(":8000"))
}
