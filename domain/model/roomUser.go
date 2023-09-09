package model

import (
    "time"

	"gorm.io/gorm"
)

type RoomUsers []RoomUser

type RoomUser struct {
	ID          int64     `json:"id"`
	RoomUserKey string    `json:"room_user_key"`
	RoomKey     string     `json:"room_key"`
	UserKey     string     `json:"user_key"`
	Host        bool      `json:"host"`
 	Status      string    `json:"status"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyRoomUser() *RoomUser {
	return &RoomUser{}
}

func (roomUser *RoomUser) IsEmpty() bool {
	return (
		roomUser.ID == 0 &&
		roomUser.RoomUserKey == "" &&
		roomUser.RoomKey == "" &&
		roomUser.UserKey == "" &&
		roomUser.Host == false &&
		roomUser.Status == "")
}
