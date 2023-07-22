package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

// user_register
type UserRegister struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

func ToUserRegister(u *model.User) *UserRegister {
	if u == nil {
		return nil
	}

	return &UserRegister{
		UserKey:  u.UserKey,
		Username: u.Username,
		Email:    u.Email,
		Message:  "user register completed",
	}
}

// email validation
type EmailValidation struct {
	Message string `json:"message"`
}

func ToEmailValidation() *EmailValidation {
	return &EmailValidation{
		Message: "email already exists",
	}
}

// user_login
type UserLogin struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

func ToUserLogin(u *model.User) *UserLogin {
	return &UserLogin{
		UserKey:  u.UserKey,
		Username: u.Username,
		Email:    u.Email,
		Token:    u.Token,
		Message:  "user login completed",
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
	UserKey  string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

func ToUserCheck(userKey string, username string, email string) *UserCheck {
	return &UserCheck{
		UserKey:  userKey,
		Username: username,
		Email:    email,
		Message:  "user check completed",
	}
}

// user_logout
type UserLogout struct {
	Message  string `json:"message"`
}

func ToUserLogout() *UserDelete {
	return &UserDelete{
		Message: "user logout completed",
	}
}
