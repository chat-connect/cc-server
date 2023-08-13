package parameter

type CreateChannelChat struct {
	Content          string  `json:"content"`
	ChannelChatImage *string `json:"channel_chat_image"`
}
