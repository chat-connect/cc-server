package service

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type GenreService interface {
	ListGenre() (genreResult *model.Genres, err error)
}

type genreService struct {
	genreRepository repository.GenreRepository
}

func NewGenreService(
		genreRepository repository.GenreRepository,
	) GenreService {
	return &genreService{
		genreRepository: genreRepository,
	}
}

// ListGenre ジャンル一覧を取得する
func (genreService *genreService) ListGenre() (genreResult *model.Genres, err error) {
	genreResult, err = genreService.genreRepository.List()
	if err != nil {
		return nil, err
	}

	return genreResult, nil
}
