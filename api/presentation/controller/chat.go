package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
)

type ChatController interface {
	ChatCreate() echo.HandlerFunc
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

// Create
// @Summary     チャット作成
// @tags        Chat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ChatCreate}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/chat_create/{roomKey} [post]
func (chatController *chatController) ChatCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")
		chatParam := &parameter.ChatCreate{}
		c.Bind(chatParam)

		chatResult, err := chatController.chatService.ChatCreate(roomKey, userKey, chatParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("chat_create", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToChatCreate(chatResult)
		response := response.SuccessWith("chat_create", 200, out)

		return c.JSON(200, response)
	}
}
