package model

import (
    "time"
)

type LinkGames []LinkGame

type LinkGame struct {
	ID            int64     `json:"id"`
	LinkGameKey   string    `json:"link_game_key"`
	ApiKey        string    `json:"api_key"`
	GameTitle     string    `json:"game_title"`
	GameImagePath string    `json:"game_image_path"`
	GameGenre     string    `json:"game_genre"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyLinkGame() *LinkGame {
	return &LinkGame{}
}

func (linkGame *LinkGame) IsEmpty() bool {
	return (
		linkGame.ID == 0 &&
		linkGame.LinkGameKey == "" &&
		linkGame.ApiKey == "" &&
		linkGame.GameTitle == "" &&
		linkGame.GameImagePath == "" &&
		linkGame.GameGenre == "")
}
