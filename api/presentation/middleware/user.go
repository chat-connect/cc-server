package middleware

import (
	"fmt"
	"strings"
	"github.com/labstack/echo/v4"

	"github.com/chat-connect/cc-server/api/service"
)

type UserMiddleware interface {
	UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type userMiddleware struct {
	userService service.UserService
}

func NewUserMiddleware(userService service.UserService) UserMiddleware {
    return &userMiddleware{
        userService: userService,
    }
}

func (userMiddleware *userMiddleware) UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")
		userKey := c.Param("userKey")

		user, err := userMiddleware.userService.FindByUserKey(userKey)
		if err != nil {
			return fmt.Errorf("Invalid user_key")
		}

		if token != user.Token {
			return fmt.Errorf("Invalid token")
		}

		return next(c)
    }
}
