package model

import (
    "time"
)

type RoomUsers []RoomUser

type RoomUser struct {
	ID          int64     `json:"id"`
	RoomUserKey string    `json:"room_user_key"`
	RoomID      int64     `json:"room_id"`
	UserID      int64     `json:"user_id"`
	Host        bool      `json:"host"`
 	Status      string    `json:"status"`
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
		roomUser.RoomID == 0 &&
		roomUser.UserID == 0 &&
		roomUser.Host == false &&
		roomUser.Status == "")
}
