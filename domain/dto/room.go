package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type RoomAndGenreAndGame struct {
	Room  model.Room
	Genre model.Genre
	Game  model.Game
}

type RoomAndGenreAndGames []RoomAndGenreAndGame
