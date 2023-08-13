package output

import (
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

type ListOpenChat struct {
	List       []ListOpenChatContent `json:"list"`
	Message    string                `json:"message"`
}

type ListOpenChatContent struct {
	OpenChatKey string    `json:"open_chat_key"`
	UserKey     string    `json:"user_key"`
	UserName    string    `json:"user_name"`
	Content     string    `json:"content"`
	ImagePath   string    `json:"image_path"`
	PostedAt    time.Time `json:"posted_at"`
}

func ToListOpenChat(c *model.OpenChats) *ListOpenChat {
	if c == nil {
		return nil
	}

	var list []ListOpenChatContent
	for _, chat := range *c {
		openChatContent := ListOpenChatContent{
			OpenChatKey: chat.OpenChatKey,
			UserKey:        chat.UserKey,
			UserName:       chat.UserName,
			Content:        chat.Content,
			ImagePath:      chat.ImagePath,
			PostedAt:       chat.PostedAt,
		}
		list = append(list, openChatContent)
	}

	return &ListOpenChat{
		List:       list,
		Message:    "list open chat created",
	}
}

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
