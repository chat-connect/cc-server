package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
)

type ChannelController interface {
	ChannelCreate() echo.HandlerFunc
}

type channelController struct {
	channelService service.ChannelService
}

func NewChannelController(
		channelService service.ChannelService,
	) ChannelController {
    return &channelController{
		channelService: channelService,
    }
}

// Create
// @Summary     チャンネル作成
// @tags        Channel
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ChannelCreate}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_create/{roomKey} [post]
func (channelController *channelController) ChannelCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")
		channelParam := &parameter.ChannelCreate{}
		c.Bind(channelParam)

		channelResult, err := channelController.channelService.ChannelCreate(roomKey, userKey, channelParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_create", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToChannelCreate(channelResult)
		response := response.SuccessWith("channel_create", 200, out)

		return c.JSON(200, response)
	}
}
