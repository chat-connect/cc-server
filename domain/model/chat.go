package model

import (
    "time"
)

type Chats []Chat

type Chat struct {
	ID        int64     `json:"id"`
	ChatKey   string    `json:"chat_key"`
	RoomKey   string    `json:"room_key"`
	UserKey   string    `json:"user_key"`
 	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyChat() *Chat {
	return &Chat{}
}

func (chat *Chat) IsEmpty() bool {
	return (
		chat.ID == 0 &&
		chat.ChatKey == "" &&
		chat.RoomKey == "" &&
		chat.UserKey == "" &&
		chat.Content == "")
}
