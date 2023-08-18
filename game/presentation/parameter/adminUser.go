package parameter

type RegisterAdminUser struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}

type LoginAdminUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminUserKey struct {
	AdminUserKey string
}
