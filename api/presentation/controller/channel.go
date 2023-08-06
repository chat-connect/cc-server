package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/api/presentation/response"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
)

type ChannelController interface {
	ChannelList() echo.HandlerFunc
	ChannelCreate() echo.HandlerFunc
	ChannelDelete() echo.HandlerFunc
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

// List
// @Summary     チャンネル一覧取得
// @tags        Channel
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ChannelList}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_list/{roomKey} [get]
func (channelController *channelController) ChannelList() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")

		channelResult, err := channelController.channelService.ChannelList(roomKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_list", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToChannelList(roomKey, channelResult)
		response := response.SuccessWith("channel_list", 200, out)

		return c.JSON(200, response)
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

// Delete
// @Summary     チャンネル削除
// @tags        Channel
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ChannelDelete}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_delete/{channelKey} [post]
func (channelController *channelController) ChannelDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")

		err := channelController.channelService.ChannelDelete(channelKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_delete", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToChannelDelete()
		response := response.SuccessWith("channel_delete", 200, out)

		return c.JSON(200, response)
	}
}
