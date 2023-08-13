package output

import (
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

type ListRoomChat struct {
	ChannelKey string                  `json:"channel_key"`
	List       []ListRoomChatContent `json:"list"`
	Message    string                   `json:"message"`
}

type ListRoomChatContent struct {
	RoomChatKey string    `json:"room_chat_key"`
	UserKey     string    `json:"user_key"`
	UserName    string    `json:"user_name"`
	Content     string    `json:"content"`
	ImagePath   string    `json:"image_path"`
	PostedAt    time.Time `json:"posted_at"`
}

func ToListRoomChat(channelKey string, c *model.RoomChats) *ListRoomChat {
	if c == nil {
		return nil
	}

	var list []ListRoomChatContent
	for _, chat := range *c {
		roomChatContent := ListRoomChatContent{
			RoomChatKey: chat.RoomChatKey,
			UserKey:        chat.UserKey,
			UserName:       chat.UserName,
			Content:        chat.Content,
			ImagePath:      chat.ImagePath,
			PostedAt:       chat.PostedAt,
		}
		list = append(list, roomChatContent)
	}

	return &ListRoomChat{
		ChannelKey: channelKey,
		List:       list,
		Message:    "list channel chat created",
	}
}

type CreateRoomChat struct {
	RoomChatKey string    `json:"room_chat_key"`
	ChannelKey  string    `json:"channel_key"`
	UserKey     string    `json:"user_key"`
	Content     string    `json:"content"`
	PostedAt    time.Time `json:"posted_at"`
	Message     string    `json:"message"`
}

func ToCreateRoomChat(c *model.RoomChat) *CreateRoomChat {
	if c == nil {
		return nil
	}

	return &CreateRoomChat{
		RoomChatKey: c.RoomChatKey,
		ChannelKey:  c.ChannelKey,
		UserKey:     c.UserKey,
		Content:     c.Content,
		PostedAt:    c.PostedAt,
		Message:     "create channel chat completed",
	}
}
