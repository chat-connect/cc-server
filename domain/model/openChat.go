package model

import (
    "time"

	"gorm.io/plugin/soft_delete"
)

type OpenChats []OpenChat

type OpenChat struct {
	ID          int64     `json:"id"`
	OpenChatKey string    `json:"open_chat_key"`
	UserKey     string    `json:"user_key"`
	UserName    string    `json:"user_name"`
 	Content     string    `json:"content"`
	ImagePath   string    `json:"image_path"`
	PostedAt    time.Time `json:"posted_at"`
	Deleted     soft_delete.DeletedAt `json:"deleted" gorm:"uniqueIndex:udx_name"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyOpenChat() *OpenChat {
	return &OpenChat{}
}

func (openChat *OpenChat) IsEmpty() bool {
	return (
		openChat.ID == 0 &&
		openChat.OpenChatKey == "" &&
		openChat.UserKey == "" &&
		openChat.UserName == "" &&
		openChat.Content == "" &&
		openChat.ImagePath == "" &&
		openChat.PostedAt.IsZero())
}
