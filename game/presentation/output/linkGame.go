package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type CreateLinkGame struct {
	LinkGameKey   string `json:"link_game_key"`
	AdminUserKey  string `json:"admin_user_key"`
	ApiKey        string `json:"api_key"`
	GameTitle     string `json:"game_title"`
	GameImagePath string `json:"game_image_path"`
	GameGenre     string `json:"game_genre"`
	Message       string `json:"message"`
}

func ToCreateLinkGame(lg *model.LinkGame) *CreateLinkGame {
	if lg == nil {
		return nil
	}

	return &CreateLinkGame{
		LinkGameKey:   lg.LinkGameKey,
		AdminUserKey:  lg.AdminUserKey,
		ApiKey:        lg.ApiKey,
		GameTitle:     lg.GameTitle,
		GameImagePath: lg.GameImagePath,
		GameGenre:     lg.GameGenre,
		Message:       "create link game completed",
	}
}
