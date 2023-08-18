package output

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
)

type ListRoom struct {
	List    []ListRoomContent `json:"list"`
	Message string            `json:"message"`
}

type ListRoomContent struct {
	RoomKey     string      `json:"room_key"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Genre       model.Genre `json:"genre"`
	Game        model.Game  `json:"game"`
}

func ToListRoom(r *dto.RoomAndGenreAndGames) *ListRoom {
	if r == nil {
		return nil
	}

	var list []ListRoomContent
	for _, room := range *r {
		roomContent := ListRoomContent{
			RoomKey:     room.Room.RoomKey,
			Name:        room.Room.Name,
			Description: room.Room.Description,
			Status:      room.Room.Status,
			Genre:       room.Genre,
			Game:        room.Game,
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
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	UserCount   int64  `json:"user_count"`
	Status      string `json:"status"`
	Genre       string `json:"genre"`
	Game        string `json:"game"`
	Message     string `json:"message"`
}

func ToCreateRoom(r *model.Room) *CreateRoom {
	if r == nil {
		return nil
	}

	return &CreateRoom{
		RoomKey:     r.RoomKey,
		Name:        r.Name,
		Description: r.Description,
		UserCount:   r.UserCount,
		Status:      r.Status,
		Genre:       r.Genre,
		Game:        r.Game,
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
