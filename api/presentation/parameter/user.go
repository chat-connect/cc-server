package parameter

type RegisterUser struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	UserImage *string `json:"user_image"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Description string  `json:"description"`
	UserImage   *string `json:"user_image"`
}

type UserKey struct {
	UserKey string
}
