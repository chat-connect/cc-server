package output

import (
	"time"
	
	"github.com/game-connect/gc-server/domain/model"
)

type CreateChat struct {
	ChatKey    string    `json:"chat_key"`
	ChannelKey string    `json:"channel_key"`
	UserKey    string    `json:"user_key"`
	Content    string    `json:"content"`
	PostedAt   time.Time `json:"posted_at"`
	Message    string    `json:"message"`
}

func ToCreateChat(c *model.Chat) *CreateChat {
	if c == nil {
		return nil
	}

	return &CreateChat{
		ChatKey:    c.ChatKey,
		ChannelKey: c.ChannelKey,
		UserKey:    c.UserKey,
		Content:    c.Content,
		PostedAt:   c.PostedAt,
		Message:    "chat create completed",
	}
}
