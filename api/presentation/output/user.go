package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type SearchUser struct {
	List    []SearchUserContent `json:"list"`
	Message string              `json:"message"`
}

type SearchUserContent struct {
	UserKey     string `json:"user_key"`
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
}

func ToSearchUser(u *model.Users) *SearchUser {
	if u == nil {
		return nil
	}

	var list []SearchUserContent
	for _, user := range *u {
		searchUserContent := SearchUserContent{
			UserKey:   user.UserKey,
			Name:      user.Name,
			ImagePath: user.ImagePath,
		}

		list = append(list, searchUserContent)
	}

	return &SearchUser{
		List:    list,
		Message: "search user created",
	}
}
