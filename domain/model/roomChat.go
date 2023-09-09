package model

import (
    "time"

	"gorm.io/plugin/soft_delete"
)

type RoomChats []RoomChat

type RoomChat struct {
	ID          int64     `json:"id"`
	RoomChatKey string    `json:"room_chat_key"`
	RoomKey     string    `json:"room_key"`
	UserKey     string    `json:"user_key"`
	UserName    string    `json:"user_name"`
 	Content     string    `json:"content"`
	ImagePath   string    `json:"image_path"`
	PostedAt    time.Time `json:"posted_at"`
	Deleted     soft_delete.DeletedAt `json:"deleted" gorm:"uniqueIndex:udx_name"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyRoomChat() *RoomChat {
	return &RoomChat{}
}

func (roomChat *RoomChat) IsEmpty() bool {
	return (
		roomChat.ID == 0 &&
		roomChat.RoomChatKey == "" &&
		roomChat.RoomKey == "" &&
		roomChat.UserKey == "" &&
		roomChat.UserName == "" &&
		roomChat.Content == "" &&
		roomChat.ImagePath == "" &&
		roomChat.PostedAt.IsZero())
}
