package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type gameDao struct {
	Conn *gorm.DB
}

func NewGameDao(conn *gorm.DB) repository.GameRepository {
	return &gameDao{
		Conn: conn,
	}
}

func (gameDao *gameDao) List() (entity *model.Games, err error) {
	entity = &model.Games{}
	res := gameDao.Conn.Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}
