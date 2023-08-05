package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// channel_list
type ChannelList struct {
	RoomKey string               `json:"room_key"`
	List    []ChannelListContent `json:"list"`
	Message string               `json:"message"`
}

type ChannelListContent struct {
	ChannelKey  string `json:"channel_key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
}

func ToChannelList(roomKey string, c *model.Channels) *ChannelList {
	if c == nil {
		return nil
	}

	var list []ChannelListContent
	for _, channel := range *c {
		channelContent := ChannelListContent{
			ChannelKey:  channel.ChannelKey,
			Type:        channel.Type,
			Name:        channel.Name,
			Explanation: channel.Explanation,
		}
		list = append(list, channelContent)
	}

	return &ChannelList{
		RoomKey: roomKey,
		List:    list,
		Message: "channel list created",
	}
}

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
