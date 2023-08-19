package output

import (
	"github.com/game-connect/gc-server/domain/dto"
)

type ListGameUser struct {
	List    []ListGameUserContent `json:"list"`
	Message string                `json:"message"`
}

type ListGameUserContent struct {
	GameKey       string `json:"game_key"`
	GameTitle     string `json:"game_title"`
	GameImagePath string `json:"game_image_path"`
}

func ToListGameUser(gu *dto.GameAndGameUsers) *ListGameUser {
	if gu == nil {
		return nil
	}

	var list []ListGameUserContent
	for _, gameUser := range *gu {
		gameUserContent := ListGameUserContent{
			GameKey:       gameUser.Game.GameKey,
			GameTitle:     gameUser.Game.GameTitle,
			GameImagePath: gameUser.Game.GameImagePath,
		}
		list = append(list, gameUserContent)
	}

	return &ListGameUser{
		List:    list,
		Message: "game list user created",
	}
}
