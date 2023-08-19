package parameter

type LoginUser struct {
	GameKey     string `json:"game_key"`
	ApiKey      string `json:"api_key"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
