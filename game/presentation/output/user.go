package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type LoginUser struct {
	UserKey   string `json:"user_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	ImagePath string `json:"image_path"`
	Description string `json:"description"`
	Message   string `json:"message"`
}

func ToLoginUser(u *model.User) *LoginUser {
	return &LoginUser{
		UserKey:   u.UserKey,
		Name:      u.Name,
		Email:     u.Email,
		Token:     u.Token,
		ImagePath: u.ImagePath,
		Description: u.Description,
		Message:   "user login completed",
	}
}
