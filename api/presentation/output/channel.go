package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// channel_create
type ChannelCreate struct {
	ChannelKey  string `json:"channel_key"`
	RoomKey     string `json:"room_key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	Message     string `json:"message"`
}

func ToChannelCreate(c *model.Channel) *ChannelCreate {
	if c == nil {
		return nil
	}

	return &ChannelCreate{
		ChannelKey:  c.ChannelKey,
		RoomKey:     c.RoomKey,
		Type:        c.Type,
		Name:        c.Name,
		Explanation: c.Explanation,
		Message:     "channel create completed",
	}
}
