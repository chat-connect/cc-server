package response

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// user_login
type UserLogin struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

func ToUserLogin(u model.User, token string) *UserLogin {
	return &UserLogin{
		UserKey:  u.UserKey,
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
		Message:  "user login completed",
	}
}

// user_register
type UserRegister struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

func ToUserRegister(u model.User) *UserRegister {
	return &UserRegister{
		UserKey:  u.UserKey,
		Username: u.Username,
		Email:    u.Email,
		Message:  "user register completed",
	}
}

// user_delete
type UserDelete struct {
	Message  string `json:"message"`
}

func ToUserDelete() *UserDelete {
	return &UserDelete{
		Message: "user delete completed",
	}
}

// user_check
type UserCheck struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserCheck(userKey string, username string, email string) *UserCheck {
	return &UserCheck{
		UserKey:  userKey,
		Username:  username,
		Email:    email,
	}
}
