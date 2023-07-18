package request

// user_register
type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// user_login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
