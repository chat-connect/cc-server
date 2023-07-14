package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/chat-connect/cc-server/log"
	"github.com/chat-connect/cc-server/infrastructure/database"
	"github.com/chat-connect/cc-server/api/controller"
	customMiddleware "github.com/chat-connect/cc-server/api/middleware"
	_ "github.com/chat-connect/cc-server/docs/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	userController := controller.NewUserController(database.NewSqlHandler())
	
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GeneratApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth
	a := e.Group("/auth")
	a.POST("/user_register", func(c echo.Context) error { return userController.Register(c) }) // auth/user_register
	a.POST("/user_login", func(c echo.Context) error { return userController.Login(c) }) // auth/user_login

	// user
	u := e.Group("/user")
	u.Use(customMiddleware.UserMiddleware)
	u.GET("/:userKey/user_check", func(c echo.Context) error { return userController.Check(c) }) // user/user_login
	u.DELETE("/:userKey/user_delete", func(c echo.Context) error { return userController.Delete(c) }) // user/user_delete

	e.Logger.Fatal(e.Start(":8000"))

	return e
}
