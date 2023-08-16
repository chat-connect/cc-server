package model

import (
    "time"
)

type Rooms []Room

type Room struct {
	ID          int64     `json:"id"`
	RoomKey     string    `json:"room_key"`
	UserKey     string    `json:"user_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImagePath   string    `json:"image_path"`
	UserCount   int64     `json:"user_count"`
 	Status      string    `json:"status"`
	Genre       string    `json:"genre"`
	Game        string    `json:"game"`
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
		room.Name == "" &&
		room.Description == "" &&
		room.ImagePath == "" &&
		room.UserCount == 0 &&
		room.Status == "" &&
		room.Genre == "" &&
		room.Game == "")
}
