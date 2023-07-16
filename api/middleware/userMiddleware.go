package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/chat-connect/cc-server/service"
	"github.com/chat-connect/cc-server/infrastructure/dao"
)

type UserMiddleware struct {
	Interactor service.UserService
}

func NewUserMiddleware(sqlHandler dao.SqlHandler) *UserMiddleware {
	return &UserMiddleware{
		Interactor: service.UserService {
				UserDao: &dao.UserDao {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (middleware *UserMiddleware) UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return fmt.Errorf("Invalid access_token")
		}
	
		token := cookie.Value
		userKey := c.Param("userKey")

		user, err := middleware.Interactor.FindByUserKey(userKey)
		if err != nil {
			return fmt.Errorf("Invalid user_key")
		}

		if token != user.Token {
			return fmt.Errorf("Invalid token")
		}

		return next(c)
    }
}
