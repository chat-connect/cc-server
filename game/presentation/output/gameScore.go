package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type UpdateGameScore struct {
	GameKey            string    `json:"game_key"`
	UserKey            string    `json:"user_key"`
	GameUsername       string    `json:"game_username"`
	GameUserImagePath  string    `json:"game_user_image_path"`
	GameScore          string    `json:"game_score"`
	GameComboScore     string    `json:"game_combo_score"`
	GameRank           string    `json:"game_rank"`
	GamePlayTime       int       `json:"game_play_time"`
	GameScoreImagePath string    `json:"game_score_image_path"`
}

func ToUpdateGameScore(gs *model.GameScore) *UpdateGameScore {
	if gs == nil {
		return nil
	}

	return &UpdateGameScore{
		GameKey:            gs.GameKey,
		UserKey:            gs.UserKey,
		GameUsername:       gs.GameUsername,
		GameUserImagePath:  gs.GameUserImagePath,
		GameScore:          gs.GameScore,
		GameComboScore:     gs.GameComboScore,
		GameRank:           gs.GameRank,
		GamePlayTime:       gs.GamePlayTime,
		GameScoreImagePath: gs.GameScoreImagePath,
	}
}