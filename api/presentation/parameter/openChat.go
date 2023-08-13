package parameter

type CreateOpenChat struct {
	Content       string  `json:"content"`
	OpenChatImage *string `json:"open_chat_image"`
}
