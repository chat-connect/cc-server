package output

import (
	"time"
	
	"github.com/game-connect/gc-server/domain/model"
)

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
		Message:     "create room chat completed",
	}
}
