package model

import (
    "time"
)

type Rooms []Room

type Room struct {
	ID          int64     `json:"id"`
	RoomKey     string    `json:"room_key"`
	UserKey     string    `json:"user_key"`
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Explanation string    `json:"explanation"`
	ImagePath   string    `json:"image_path"`
	UserCount   int64     `json:"user_count"`
 	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyRoom() *Room {
	return &Room{}
}

func (room *Room) IsEmpty() bool {
	return (
		room.ID == 0 &&
		room.RoomKey == "" &&
		room.UserKey == "" &&
		room.UserID == 0 &&
		room.Name == "" &&
		room.Explanation == "" &&
		room.ImagePath == "" &&
		room.UserCount == 0 &&
		room.Status == "")
}
