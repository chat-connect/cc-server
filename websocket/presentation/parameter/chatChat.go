package parameter

type CreateChannelChat struct {
	UserKey    string `json:"user_key"`
	Token      string `json:"token"`
	Content    string `json:"content"`
	ChannelChatImage *string `json:"channel_chat_image"`
}
