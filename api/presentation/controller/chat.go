package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type ChatController interface {
	ListChat() echo.HandlerFunc
	CreateChat() echo.HandlerFunc
}

type chatController struct {
	chatService service.ChatService
}

func NewChatController(
		chatService service.ChatService,
	) ChatController {
    return &chatController{
		chatService: chatService,
    }
}

// List
// @Summary     チャット一覧取得
// @tags        Chat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/chat_list/{channelKey} [get]
func (chatController *chatController) ListChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")

		chatResult, err := chatController.chatService.ListChat(channelKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("chat_list", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListChat(channelKey, chatResult)
		response := response.SuccessWith("chat_list", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     チャット作成
// @tags        Chat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/chat_create/{channelKey} [post]
func (chatController *chatController) CreateChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")
		userKey := c.Param("userKey")
		chatParam := &parameter.CreateChat{}
		c.Bind(chatParam)

		chatResult, err := chatController.chatService.CreateChat(channelKey, userKey, chatParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("chat_create", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateChat(chatResult)
		response := response.SuccessWith("chat_create", 200, out)

		return c.JSON(200, response)
	}
}
