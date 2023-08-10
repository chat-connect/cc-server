package output

import (
	"github.com/chat-connect/cc-server/domain/model"
)

type RegisterUser struct {
	UserKey string `json:"user_key"`
	Name    string `json:"name"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

func ToRegisterUser(u *model.User) *RegisterUser {
	if u == nil {
		return nil
	}

	return &RegisterUser{
		UserKey: u.UserKey,
		Name:    u.Name,
		Email:   u.Email,
		Message: "user register completed",
	}
}

type EmailValidation struct {
	Message string `json:"message"`
}

func ToEmailValidation() *EmailValidation {
	return &EmailValidation{
		Message: "email already exists",
	}
}

type LoginUser struct {
	UserKey string `json:"user_key"`
	Name    string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Message  string `json:"message"`
}

func ToLoginUser(u *model.User) *LoginUser {
	return &LoginUser{
		UserKey: u.UserKey,
		Name:    u.Name,
		Email:   u.Email,
		Token:   u.Token,
		Message: "user login completed",
	}
}

type DeleteUser struct {
	Message  string `json:"message"`
}

func ToDeleteUser() *DeleteUser {
	return &DeleteUser{
		Message: "user delete completed",
	}
}

type CheckUser struct {
	UserKey string `json:"user_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func ToCheckUser(userKey string, name string, email string) *CheckUser {
	return &CheckUser{
		UserKey: userKey,
		Name:    name,
		Email:   email,
		Message: "user check completed",
	}
}

type LogoutUser struct {
	Message  string `json:"message"`
}

func ToLogoutUser() *DeleteUser {
	return &DeleteUser{
		Message: "user logout completed",
	}
}
