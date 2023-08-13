package output

import (
	"time"
	
	"github.com/game-connect/gc-server/domain/model"
)

type CreateChannelChat struct {
	ChannelChatKey string    `json:"channel_chat_key"`
	ChannelKey     string    `json:"channel_key"`
	UserKey        string    `json:"user_key"`
	Content        string    `json:"content"`
	PostedAt       time.Time `json:"posted_at"`
	Message        string    `json:"message"`
}

func ToCreateChannelChat(c *model.ChannelChat) *CreateChannelChat {
	if c == nil {
		return nil
	}

	return &CreateChannelChat{
		ChannelChatKey: c.ChannelChatKey,
		ChannelKey:     c.ChannelKey,
		UserKey:        c.UserKey,
		Content:        c.Content,
		PostedAt:       c.PostedAt,
		Message:        "create channel chat completed",
	}
}
