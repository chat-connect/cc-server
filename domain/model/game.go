package model

import (
    "time"
)

type Games []Game

type Game struct {
	ID            int64     `json:"id"`
	GameKey       string    `json:"game_key"`
	GenreKey      string    `json:"genre_key"`
	AdminUserKey  string    `json:"admin_user_key"`
	ApiKey        string    `json:"api_key"`
	GameTitle     string    `json:"game_title"`
	GameImagePath string    `json:"game_image_path"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGame() *Game {
	return &Game{}
}

func (game *Game) IsEmpty() bool {
	return (
		game.ID == 0 &&
		game.GameKey == "" &&
		game.GenreKey == "" &&
		game.AdminUserKey == "" &&
		game.ApiKey == "" &&
		game.GameTitle == "" &&
		game.GameImagePath == "")
}

func (games *Games) SearchGameKey(gameKey string) *Game {
	for _, game := range *games {
		if game.GameKey == gameKey {
			return &game
		}
	}
	
	return nil
}
