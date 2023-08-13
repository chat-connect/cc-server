package parameter

type CreateOpenChat struct {
	UserKey       string `json:"user_key"`
	Token         string `json:"token"`
	Content       string `json:"content"`
	OpenChatImage *string `json:"open_chat_image"`
}
