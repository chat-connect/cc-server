package parameter

type CreateChat struct {
	UserKey string `json:"user_key"`
	Token   string `json:"token"`
	Content string `json:"content"`
}
