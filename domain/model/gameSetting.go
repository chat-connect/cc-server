package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type GameSetting struct {
	ID                 int64      `json:"id"`
	GameKey            string     `json:"game_key"`
	AdminUserKey       string     `json:"admin_user_key"`
	GameScore          bool       `json:"game_score"`
	GameComboScore     bool       `json:"game_combo_score"`
	GameRank           bool       `json:"game_rank"`
	GamePlayTime       bool       `json:"game_play_time"`
	GameScoreImagePath bool       `json:"game_score_image_path"`
	Deleted            soft_delete.DeletedAt `json:"deleted" gorm:"uniqueIndex:udx_name"`
	CreatedAt          time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGameSetting() *GameSetting {
	return &GameSetting{}
}

func (gameSetting *GameSetting) IsEmpty() bool {
	return (
		gameSetting.ID == 0 &&
		gameSetting.GameKey == "" &&
		gameSetting.AdminUserKey == "" &&
		gameSetting.GameScore == false &&
		gameSetting.GameComboScore == false &&
		gameSetting.GameRank == false &&
		gameSetting.GamePlayTime == false &&
		gameSetting.GameScoreImagePath == false)
}
