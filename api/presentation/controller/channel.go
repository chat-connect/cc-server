package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
	"github.com/game-connect/gc-server/api/presentation/parameter"
)

type ChannelController interface {
	ListChannel() echo.HandlerFunc
	CreateChannel() echo.HandlerFunc
	DeleteChannel() echo.HandlerFunc
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
// @Success     200  {object} response.Success{items=output.ListChannel}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_list/{roomKey} [get]
func (channelController *channelController) ListChannel() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")

		channelResult, err := channelController.channelService.ListChannel(roomKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_list", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListChannel(roomKey, channelResult)
		response := response.SuccessWith("channel_list", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     チャンネル作成
// @tags        Channel
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.CreateChannel}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_create/{roomKey} [post]
func (channelController *channelController) CreateChannel() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")
		channelParam := &parameter.CreateChannel{}
		c.Bind(channelParam)

		channelResult, err := channelController.channelService.CreateChannel(roomKey, userKey, channelParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_create", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateChannel(channelResult)
		response := response.SuccessWith("channel_create", 200, out)

		return c.JSON(200, response)
	}
}

// Delete
// @Summary     チャンネル削除
// @tags        Channel
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.DeleteChannel}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /channel/{userKey}/channel_delete/{channelKey} [post]
func (channelController *channelController) DeleteChannel() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		channelKey := c.Param("channelKey")

		err := channelController.channelService.DeleteChannel(channelKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("channel_delete", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToDeleteChannel()
		response := response.SuccessWith("channel_delete", 200, out)

		return c.JSON(200, response)
	}
}
