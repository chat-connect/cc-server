package repository

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GenreRepository interface {
	FindByGenreKey(genreKey string) (entity *model.Genre, err error)
	List() (entity *model.Genres, err error)
}
