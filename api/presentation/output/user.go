package output

import (
	"github.com/game-connect/gc-server/domain/dto"
)

type SearchUser struct {
	List    []SearchUserContent `json:"list"`
	Message string              `json:"message"`
}

type SearchUserContent struct {
	UserKey     string `json:"user_key"`
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
	Following   bool   `json:"following"`
}

func ToSearchUser(u *dto.SearchUsers) *SearchUser {
	if u == nil {
		return nil
	}

	var list []SearchUserContent
	for _, user := range *u {
		searchUserContent := SearchUserContent{
			UserKey:   user.User.UserKey,
			Name:      user.User.Name,
			ImagePath: user.User.ImagePath,
			Following: user.Following,
		}

		list = append(list, searchUserContent)
	}

	return &SearchUser{
		List:    list,
		Message: "search user created",
	}
}
