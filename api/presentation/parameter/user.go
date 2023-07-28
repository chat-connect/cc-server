package parameter

// user_register
type UserRegister struct {
	Name string     `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// user_login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// user_key
type UserKey struct {
	UserKey string
}
