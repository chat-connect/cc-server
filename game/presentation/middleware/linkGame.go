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

type LinkGameMiddleware interface {
	CheckApiKey(next echo.HandlerFunc) echo.HandlerFunc
}

type linkGameMiddleware struct {
	linkGameService service.LinkGameService
}

func NewLinkGameMiddleware(
		linkGameService service.LinkGameService,
	) LinkGameMiddleware {
    return &linkGameMiddleware{
		linkGameService: linkGameService,
    }
}

func (linkGameMiddleware *linkGameMiddleware) CheckApiKey(next echo.HandlerFunc) echo.HandlerFunc {
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

        _, err = linkGameMiddleware.linkGameService.FindByApiKey(userParam.ApiKey)
        if err != nil {
            return fmt.Errorf("Invalid api key")
        }

        return next(c)
    }
}