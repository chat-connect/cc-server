package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type ListChat struct {
	ChannelKey string            `json:"channel_key"`
	List       []ListChatContent `json:"list"`
	Message    string            `json:"message"`
}

type ListChatContent struct {
	ChatKey string  `json:"chat_key"`
	UserKey string  `json:"user_key"`
	UserName string `json:"user_name"`
	Content string  `json:"content"`
}

func ToListChat(channelKey string, c *model.Chats) *ListChat {
	if c == nil {
		return nil
	}

	var list []ListChatContent
	for _, chat := range *c {
		chatContent := ListChatContent{
			ChatKey: chat.ChatKey,
			UserKey: chat.UserKey,
			UserName: chat.UserName,
			Content: chat.Content,
		}
		list = append(list, chatContent)
	}

	return &ListChat{
		ChannelKey: channelKey,
		List:       list,
		Message:    "chat list created",
	}
}

type CreateChat struct {
	ChatKey    string `json:"chat_key"`
	ChannelKey string `json:"channel_key"`
	UserKey    string `json:"user_key"`
	Content    string `json:"content"`
	Message    string `json:"message"`
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
		Message:    "chat create completed",
	}
}
