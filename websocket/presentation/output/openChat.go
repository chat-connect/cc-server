package output

import (
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

type CreateOpenChat struct {
	OpenChatKey string    `json:"open_chat_key"`
	UserKey     string    `json:"user_key"`
	Content     string    `json:"content"`
	PostedAt    time.Time `json:"posted_at"`
	Message     string    `json:"message"`
}

func ToCreateOpenChat(c *model.OpenChat) *CreateOpenChat {
	if c == nil {
		return nil
	}

	return &CreateOpenChat{
		OpenChatKey: c.OpenChatKey,
		UserKey:     c.UserKey,
		Content:     c.Content,
		PostedAt:    c.PostedAt,
		Message:     "create open chat completed",
	}
}
