package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type ChannelChatController interface {
	ListChannelChat() echo.HandlerFunc
	CreateChannelChat() echo.HandlerFunc
}

type channelChatController struct {
	channelChatService service.ChannelChatService
}

func NewChannelChatController(
		channelChatService service.ChannelChatService,
	) ChannelChatController {
    return &channelChatController{
		channelChatService: channelChatService,
    }
}

// List
// @Summary     チャット一覧取得
// @tags        ChannelChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListChannelChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/list_channel_chat/{channelKey} [get]
func (channelChatController *channelChatController) ListChannelChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")

		channelChatResult, err := channelChatController.channelChatService.ListChannelChat(channelKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_channel_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListChannelChat(channelKey, channelChatResult)
		response := response.SuccessWith("list_channel_chat", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     チャット作成
// @tags        ChannelChat
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateChannelChat}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /chat/{userKey}/create_channel_chat/{channelKey} [post]
func (channelChatController *channelChatController) CreateChannelChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")
		userKey := c.Param("userKey")
		channelChatParam := &parameter.CreateChannelChat{}
		c.Bind(channelChatParam)

		channelChatResult, err := channelChatController.channelChatService.CreateChannelChat(channelKey, userKey, channelChatParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_channel_chat", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateChannelChat(channelChatResult)
		response := response.SuccessWith("create_channel_chat", 200, out)

		return c.JSON(200, response)
	}
}
