package model

import (
    "time"
)

type Chats []Chat

type Chat struct {
	ID         int64     `json:"id"`
	ChatKey    string    `json:"chat_key"`
	ChannelKey string    `json:"channel_key"`
	UserKey    string    `json:"user_key"`
	UserName   string    `json:"user_name"`
 	Content    string    `json:"content"`
	PostedAt   time.Time `json:"posted_at"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyChat() *Chat {
	return &Chat{}
}

func (chat *Chat) IsEmpty() bool {
	return (
		chat.ID == 0 &&
		chat.ChatKey == "" &&
		chat.ChannelKey == "" &&
		chat.UserKey == "" &&
		chat.UserName == "" &&
		chat.Content == "" &&
		chat.PostedAt.IsZero())
}
