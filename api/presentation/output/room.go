package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// room_list
type RoomList struct {
	List    []RoomListContent `json:"list"`
	Message string            `json:"message"`
}

type RoomListContent struct {
	RoomKey     string  `json:"chat_key"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
}

func ToRoomList(r *model.Rooms) *RoomList {
	if r == nil {
		return nil
	}

	var list []RoomListContent
	for _, room := range *r {
		roomContent := RoomListContent{
			RoomKey:     room.RoomKey,
			Name:        room.Name,
			Explanation: room.Explanation,
		}
		list = append(list, roomContent)
	}

	return &RoomList{
		List:    list,
		Message: "room list created",
	}
}

// room_create
type RoomCreate struct {
	RoomKey     string `json:"room_key"`
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
		Name:        r.Name,
		Explanation: r.Explanation,
		UserCount:   r.UserCount,
		Status:      r.Status,
		Message: "room create completed",
	}
}

// room_join
type RoomJoin struct {
	RoomUserKey string `json:"room_user_key"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

func ToRoomJoin(ru *model.RoomUser) *RoomJoin {
	if ru == nil {
		return nil
	}

	return &RoomJoin{
		RoomUserKey: ru.RoomUserKey,
		Status:      ru.Status,
		Message: "room join completed",
	}
}

// room_out
type RoomOut struct {
	Message string `json:"message"`
}

func ToRoomOut() *RoomOut {
	return &RoomOut{
		Message: "room out completed",
	}
}

// room_delete
type RoomDelete struct {
	Message string `json:"message"`
}

func ToRoomDelete() *RoomOut {
	return &RoomOut{
		Message: "room delete completed",
	}
}
