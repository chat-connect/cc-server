package model

import (
    "time"

	"gorm.io/plugin/soft_delete"
)

type ChannelChats []ChannelChat

type ChannelChat struct {
	ID             int64     `json:"id"`
	ChannelChatKey string    `json:"channel_chat_key"`
	ChannelKey     string    `json:"channel_key"`
	UserKey        string    `json:"user_key"`
	UserName       string    `json:"user_name"`
 	Content        string    `json:"content"`
	ImagePath      string    `json:"image_path"`
	PostedAt       time.Time `json:"posted_at"`
	Deleted        soft_delete.DeletedAt `json:"deleted" gorm:"uniqueIndex:udx_name"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyChannelChat() *ChannelChat {
	return &ChannelChat{}
}

func (channelChat *ChannelChat) IsEmpty() bool {
	return (
		channelChat.ID == 0 &&
		channelChat.ChannelChatKey == "" &&
		channelChat.ChannelKey == "" &&
		channelChat.UserKey == "" &&
		channelChat.UserName == "" &&
		channelChat.Content == "" &&
		channelChat.ImagePath == "" &&
		channelChat.PostedAt.IsZero())
}
