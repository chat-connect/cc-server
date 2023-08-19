package parameter

type LoginUser struct {
	LinkGameKey string `json:"link_game_key"`
	ApiKey   string `json:"api_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
