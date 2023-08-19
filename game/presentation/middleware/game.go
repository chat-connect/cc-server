package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/labstack/echo/v4"

	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type GameMiddleware interface {
	CheckApiKey(next echo.HandlerFunc) echo.HandlerFunc
}

type gameMiddleware struct {
	gameService service.GameService
}

func NewGameMiddleware(
		gameService service.GameService,
	) GameMiddleware {
    return &gameMiddleware{
		gameService: gameService,
    }
}

func (gameMiddleware *gameMiddleware) CheckApiKey(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        bodyBytes, err := ioutil.ReadAll(c.Request().Body)
        if err != nil {
            return err
        }

        userParam := &parameter.LoginUser{}
        if err := json.Unmarshal(bodyBytes, userParam); err != nil {
            return err
        }

        c.Request().Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

        _, err = gameMiddleware.gameService.FindByApiKey(userParam.ApiKey)
        if err != nil {
            return fmt.Errorf("Invalid api key")
        }

        return next(c)
    }
}