package output

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
)

type GetUser struct {
	UserKey     string `json:"user_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ImagePath   string `json:"image_path"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func ToGetUser(u *model.User) *GetUser {
	return &GetUser{
		UserKey:     u.UserKey,
		Name:        u.Name,
		Email:       u.Email,
		Token:       u.Token,
		ImagePath:   u.ImagePath,
		Description: u.Description,
		Message:     "get user completed",
	}
}

func ToUpdateUser(u *model.User) *GetUser {
	return &GetUser{
		UserKey:     u.UserKey,
		Name:        u.Name,
		Email:       u.Email,
		Token:       u.Token,
		ImagePath:   u.ImagePath,
		Description: u.Description,
		Message:     "update user completed",
	}
}

type SearchUser struct {
	List    []SearchUserContent `json:"list"`
	Message string              `json:"message"`
}

type SearchUserContent struct {
	UserKey     string `json:"user_key"`
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
	Following   bool   `json:"following"`
}

func ToSearchUser(u *dto.SearchUsers) *SearchUser {
	if u == nil {
		return nil
	}

	var list []SearchUserContent
	for _, user := range *u {
		searchUserContent := SearchUserContent{
			UserKey:     user.User.UserKey,
			Name:        user.User.Name,
			ImagePath:   user.User.ImagePath,
			Description: user.User.Description,
			Following:   user.Following,
		}

		list = append(list, searchUserContent)
	}

	return &SearchUser{
		List:    list,
		Message: "search user created",
	}
}
