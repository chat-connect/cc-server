package parameter

// channel_create
type ChannelCreate struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
}
