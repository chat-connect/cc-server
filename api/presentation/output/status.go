package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type UserStatus struct {
	UserKey   string `json:"user_key"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

func ToUserStatus(u *model.User) *UserStatus {
	return &UserStatus{
		UserKey:   u.UserKey,
		Name:      u.Name,
		ImagePath: u.ImagePath,
	}
}
