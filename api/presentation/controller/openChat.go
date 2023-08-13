package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type OpenChatController interface {
	ListOpenChat() echo.HandlerFunc
	CreateOpenChat() echo.HandlerFunc
}

type openChatController struct {
	openChatService service.OpenChatService
}

func NewOpenChatController(
		openChatService service.OpenChatService,
	) OpenChatController {
    return &openChatController{
		openChatService: openChatService,
    }
}

// List
// @Summary     チャット一覧取得
// @tags        OpenChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListOpenChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/list_open_chat [get]
func (openChatController *openChatController) ListOpenChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		openChatResult, err := openChatController.openChatService.ListOpenChat()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_open_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListOpenChat(openChatResult)
		response := response.SuccessWith("list_open_chat", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     チャット作成
// @tags        OpenChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateOpenChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/create_open_chat/{channelKey} [post]
func (openChatController *openChatController) CreateOpenChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		openChatParam := &parameter.CreateOpenChat{}
		c.Bind(openChatParam)

		openChatResult, err := openChatController.openChatService.CreateOpenChat(userKey, openChatParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_open_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateOpenChat(openChatResult)
		response := response.SuccessWith("create_open_chat", 200, out)

		return c.JSON(200, response)
	}
}
