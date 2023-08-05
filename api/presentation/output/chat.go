package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// chat_list
type ChatList struct {
	ChannelKey string         `json:"room_key"`
	List    []ChatListContent `json:"list"`
	Message string            `json:"message"`
}

type ChatListContent struct {
	ChatKey string  `json:"chat_key"`
	UserKey string  `json:"user_key"`
	UserName string `json:"user_name"`
	Content string  `json:"content"`
}

func ToChatList(channelKey string, c *model.Chats) *ChatList {
	if c == nil {
		return nil
	}

	var list []ChatListContent
	for _, chat := range *c {
		chatContent := ChatListContent{
			ChatKey: chat.ChatKey,
			UserKey: chat.UserKey,
			UserName: chat.UserName,
			Content: chat.Content,
		}
		list = append(list, chatContent)
	}

	return &ChatList{
		ChannelKey: channelKey,
		List:       list,
		Message:    "chat list created",
	}
}

// chat_create
type ChatCreate struct {
	ChatKey    string `json:"chat_key"`
	ChannelKey string `json:"channel_key"`
	UserKey    string `json:"user_key"`
	Content    string `json:"content"`
	Message    string `json:"message"`
}

func ToChatCreate(c *model.Chat) *ChatCreate {
	if c == nil {
		return nil
	}

	return &ChatCreate{
		ChatKey:    c.ChatKey,
		ChannelKey: c.ChannelKey,
		UserKey:    c.UserKey,
		Content:    c.Content,
		Message:    "chat create completed",
	}
}
