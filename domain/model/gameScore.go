package model

import (
    "time"
)

type GameScores []GameScore

type GameScore struct {
	ID                 int64     `json:"id"`
	GameScoreKey       string    `json:"game_score_key"`
	GameKey            string    `json:"game_key"`
	UserKey            string    `json:"user_key"`
	GameScore          int       `json:"game_score"`
	GameComboScore     int       `json:"game_combo_score"`
	GameRank           int       `json:"game_rank"`
	GamePlayTime       int       `json:"game_play_time"`
	GameScoreImagePath string    `json:"game_score_image_path"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGameScore() *GameScore {
	return &GameScore{}
}

func (gameScore *GameScore) IsEmpty() bool {
	return (
		gameScore.ID == 0 &&
		gameScore.GameScoreKey == "" &&
		gameScore.GameKey == "" &&
		gameScore.UserKey == "" &&
		gameScore.GameScore == 0 &&
		gameScore.GameComboScore == 0 &&
		gameScore.GameRank == 0 &&
		gameScore.GamePlayTime == 0 &&
		gameScore.GameScoreImagePath == "")
}

func (scores GameScores) Reverse() GameScores {
    reversed := make(GameScores, len(scores))
    for i, j := 0, len(scores)-1; i < len(scores); i, j = i+1, j-1 {
        reversed[i] = scores[j]
    }
    return reversed
}
