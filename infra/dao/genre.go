package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type genreDao struct {
	Conn *gorm.DB
}

func NewGenreDao(conn *gorm.DB) repository.GenreRepository {
	return &genreDao{
		Conn: conn,
	}
}

func (genreDao *genreDao) FindByGenreKey(genreKey string) (entity *model.Genre, err error) {
	entity = &model.Genre{}
	res := genreDao.Conn.
		Where("genre_key = ?", genreKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (genreDao *genreDao) List() (entity *model.Genres, err error) {
	entity = &model.Genres{}
	res := genreDao.Conn.Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}
