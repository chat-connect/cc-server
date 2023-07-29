package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// user_register
type RoomCreate struct {
	RoomKey     string `json:"room_key"`
	UserKey     string `json:"user_key"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	ImagePath   string `json:"image_path"`
	UserCount   int64  `json:"user_count"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

func ToRoomCreate(r *model.Room) *RoomCreate {
	if r == nil {
		return nil
	}

	return &RoomCreate{
		RoomKey:     r.RoomKey,
		UserKey:     r.UserKey,
		Name:        r.Name,
		Explanation: r.Explanation,
		UserCount:   r.UserCount,
		Status:      r.Status,
		Message: "room create completed",
	}
}
