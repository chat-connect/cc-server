package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type ListChannel struct {
	RoomKey string               `json:"room_key"`
	List    []ListChannelContent `json:"list"`
	Message string               `json:"message"`
}

type ListChannelContent struct {
	ChannelKey  string `json:"channel_key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
}

func ToListChannel(roomKey string, c *model.Channels) *ListChannel {
	if c == nil {
		return nil
	}

	var list []ListChannelContent
	for _, channel := range *c {
		channelContent := ListChannelContent{
			ChannelKey:  channel.ChannelKey,
			Type:        channel.Type,
			Name:        channel.Name,
			Explanation: channel.Explanation,
		}
		list = append(list, channelContent)
	}

	return &ListChannel{
		RoomKey: roomKey,
		List:    list,
		Message: "channel list created",
	}
}

type CreateChannel struct {
	ChannelKey  string `json:"channel_key"`
	RoomKey     string `json:"room_key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	Message     string `json:"message"`
}

func ToCreateChannel(c *model.Channel) *CreateChannel {
	if c == nil {
		return nil
	}

	return &CreateChannel{
		ChannelKey:  c.ChannelKey,
		RoomKey:     c.RoomKey,
		Type:        c.Type,
		Name:        c.Name,
		Explanation: c.Explanation,
		Message:     "channel create completed",
	}
}

type DeleteChannel struct {
	Message  string `json:"message"`
}

func ToDeleteChannel() *DeleteChannel {
	return &DeleteChannel{
		Message: "channel delete completed",
	}
}
