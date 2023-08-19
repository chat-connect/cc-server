package output

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
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

type ListGameScore struct {
	GameKey           string                 `json:"game_key"`
	GameTitle         string                 `json:"game_title"`
	GameImagePath     string                 `json:"game_image_path"`
	List              []ListGameScoreContent `json:"list"`
	Message           string                 `json:"message"`
}

type ListGameScoreContent struct {
	GameScore          string `json:"game_score"`
	GameComboScore     string `json:"game_combo_score"`
	GameRank           string `json:"game_rank"`
	GamePlayTime       int    `json:"game_play_time"`
	GameScoreImagePath string `json:"game_score_image_path"`
}

func ToListGameScore(gs *dto.GameAndGameScore) *ListGameScore {
	if gs == nil {
		return nil
	}

	var list []ListGameScoreContent
	for _, gameScore := range gs.GameScores {
		gameScoreContent := ListGameScoreContent{
			GameScore:          gameScore.GameScore,
			GameComboScore:     gameScore.GameComboScore,
			GameRank:           gameScore.GameRank,
			GamePlayTime:       gameScore.GamePlayTime,
			GameScoreImagePath: gameScore.GameScoreImagePath,
		}
		list = append(list, gameScoreContent)
	}

	return &ListGameScore{
		GameKey:       gs.Game.GameKey,
		GameTitle:     gs.Game.GameTitle,
		GameImagePath: gs.Game.GameTitle,
		List:          list,
		Message:       "game list score created",
	}
}
