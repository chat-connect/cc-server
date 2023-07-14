package middleware

import (
	"fmt"
	
	"github.com/labstack/echo/v4"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		baseToken := c.Request().Header.Get("Authorization")
		if baseToken == "" {
			return fmt.Errorf("Invalid token")
		}

		userKey := c.Param("userKey")
		fmt.Println(userKey)
		if userKey == "" {
			return fmt.Errorf("Invalid user_key")
		}

		return next(c)
    }
}
