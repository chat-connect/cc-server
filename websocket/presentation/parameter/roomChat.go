package parameter

type CreateRoomChat struct {
	UserKey       string `json:"user_key"`
	Token         string `json:"token"`
	Content       string `json:"content"`
	RoomChatImage *string `json:"room_chat_image"`
}
