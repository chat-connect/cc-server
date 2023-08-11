package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

type ListRoom struct {
	List    []ListRoomContent `json:"list"`
	Message string            `json:"message"`
}

type ListRoomContent struct {
	RoomKey     string `json:"room_key"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	Status      string `json:"status"`
}

func ToListRoom(r *model.Rooms) *ListRoom {
	if r == nil {
		return nil
	}

	var list []ListRoomContent
	for _, room := range *r {
		roomContent := ListRoomContent{
			RoomKey:     room.RoomKey,
			Name:        room.Name,
			Explanation: room.Explanation,
			Status:      room.Status,
		}
		list = append(list, roomContent)
	}

	return &ListRoom{
		List:    list,
		Message: "room list created",
	}
}

type CreateRoom struct {
	RoomKey     string `json:"room_key"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	ImagePath   string `json:"image_path"`
	UserCount   int64  `json:"user_count"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

func ToCreateRoom(r *model.Room) *CreateRoom {
	if r == nil {
		return nil
	}

	return &CreateRoom{
		RoomKey:     r.RoomKey,
		Name:        r.Name,
		Explanation: r.Explanation,
		UserCount:   r.UserCount,
		Status:      r.Status,
		Message: "room create completed",
	}
}

type JoinRoom struct {
	RoomUserKey string `json:"room_user_key"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

func ToJoinRoom(ru *model.RoomUser) *JoinRoom {
	if ru == nil {
		return nil
	}

	return &JoinRoom{
		RoomUserKey: ru.RoomUserKey,
		Status:      ru.Status,
		Message: "room join completed",
	}
}

type OutRoom struct {
	Message string `json:"message"`
}

func ToOutRoom() *OutRoom {
	return &OutRoom{
		Message: "room out completed",
	}
}

type DeleteRoom struct {
	Message string `json:"message"`
}

func ToDeleteRoom() *OutRoom {
	return &OutRoom{
		Message: "room delete completed",
	}
}
