package model

import (
    "time"
)

type Games []Game

type Game struct {
	ID          int64     `json:"id"`
	GameKey     string    `json:"game_key"`
	GenreKey    string    `json:"genre_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGame() *Game {
	return &Game{}
}

func (game *Game) IsEmpty() bool {
	return (
		game.ID == 0 &&
		game.GameKey == "" &&
		game.GenreKey == "" &&
		game.Name == "" &&
		game.Description == "" &&
		game.Type == "")
}
