package model

import (
    "time"

	"gorm.io/gorm"
)

type Channels []Channel

type Channel struct {
	ID          int64     `json:"id"`
	ChannelKey  string    `json:"channel_key"`
	RoomKey     string    `json:"room_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyChannel() *Channel {
	return &Channel{}
}

func (channel *Channel) IsEmpty() bool {
	return (
		channel.ID == 0 &&
		channel.ChannelKey == "" &&
		channel.RoomKey == "" &&
		channel.Description == "" &&
		channel.Type == "")
}
