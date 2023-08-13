package parameter

type CreateRoomChat struct {
	Content       string  `json:"content"`
	RoomChatImage *string `json:"room_chat_image"`
}
