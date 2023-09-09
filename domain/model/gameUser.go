package model

import (
    "time"

	"gorm.io/gorm"
)

type GameUsers []GameUser

type GameUser struct {
	ID          int64   `json:"id"`
	GameUserKey string  `json:"game_user_key"`
	UserKey     string  `json:"user_key"`
	GameKey     string  `json:"link_game_key"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGameUser() *GameUser {
	return &GameUser{}
}

func (gameUser *GameUser) IsEmpty() bool {
	return (
		gameUser.ID == 0 &&
		gameUser.GameUserKey == "" &&
		gameUser.UserKey == "" &&
		gameUser.GameKey == "")
}
