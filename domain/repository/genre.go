package repository

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GenreRepository interface {
	List() (entity *model.Genres, err error)
}
