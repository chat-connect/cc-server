package output

import (
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

type ListChannelChat struct {
	ChannelKey string                  `json:"channel_key"`
	List       []ListChannelChatContent `json:"list"`
	Message    string                   `json:"message"`
}

type ListChannelChatContent struct {
	ChannelChatKey  string    `json:"channel_chat_key"`
	UserKey         string    `json:"user_key"`
	UserName        string    `json:"user_name"`
	Content         string    `json:"content"`
	ImagePath       string    `json:"image_path"`
	PostedAt        time.Time `json:"posted_at"`
}

func ToListChannelChat(channelKey string, c *model.ChannelChats) *ListChannelChat {
	if c == nil {
		return nil
	}

	var list []ListChannelChatContent
	for _, chat := range *c {
		channelChatContent := ListChannelChatContent{
			ChannelChatKey: chat.ChannelChatKey,
			UserKey:        chat.UserKey,
			UserName:       chat.UserName,
			Content:        chat.Content,
			ImagePath:      chat.ImagePath,
			PostedAt:       chat.PostedAt,
		}
		list = append(list, channelChatContent)
	}

	return &ListChannelChat{
		ChannelKey: channelKey,
		List:       list,
		Message:    "list channel chat created",
	}
}

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
