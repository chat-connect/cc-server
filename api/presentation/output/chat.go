package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// chat_create
type ChatCreate struct {
	ChatKey string `json:"chat_key"`
	RoomKey string `json:"room_key"`
	UserKey string `json:"user_key"`
	Content string `json:"content"`
	Message string `json:"message"`
}

func ToChatCreate(c *model.Chat) *ChatCreate {
	if c == nil {
		return nil
	}

	return &ChatCreate{
		ChatKey: c.ChatKey,
		RoomKey: c.RoomKey,
		UserKey: c.UserKey,
		Content: c.Content,
		Message: "chat create completed",
	}
}
