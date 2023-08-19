package parameter

type LoginUser struct {
	ApiKey   string `json:"api_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
