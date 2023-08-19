package model

import (
    "time"
)

type GameScores []GameScore

type GameScore struct {
	ID                 int64     `json:"id"`
	GameScoreKey       string    `json:"game_score_key"`
	LinkGameKey        string    `json:"link_game_key"`
	UserKey            string    `json:"user_key"`
	GameUsername       string    `json:"game_username"`
	GameUserImagePath  string    `json:"game_user_image_path"`
	GameScore          string    `json:"game_score"`
	GameComboScore     string    `json:"game_combo_score"`
	GameRank           string    `json:"game_rank"`
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
		gameScore.LinkGameKey == "" &&
		gameScore.UserKey == "" &&
		gameScore.GameUsername == "" &&
		gameScore.GameUserImagePath == "" &&
		gameScore.GameScore == "" &&
		gameScore.GameComboScore == "" &&
		gameScore.GameRank == "" &&
		gameScore.GamePlayTime == 0 &&
		gameScore.GameScoreImagePath == "")
}
