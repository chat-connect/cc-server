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

func (gameDao *gameDao) FindByGameKey(gameKey string) (entity *model.Game, err error) {
	entity = &model.Game{}
	res := gameDao.Conn.
		Where("game_key = ?", gameKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (gameDao *gameDao) FindByApiKey(apiKey string) (entity *model.Game, err error) {
	entity = &model.Game{}
	res := gameDao.Conn.Where("api_key = ?", apiKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (gameDao *gameDao) List() (entity *model.Games, err error) {
	entity = &model.Games{}
	res := gameDao.Conn.Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (gameDao *gameDao) Insert(gameModel *model.Game, tx *gorm.DB) (entity *model.Game, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = gameDao.Conn
	}

	entity = &model.Game{
		GameKey:       gameModel.GameKey,
		GenreKey:      gameModel.GenreKey,
		AdminUserKey:  gameModel.AdminUserKey,
		ApiKey:        gameModel.ApiKey,
		GameTitle:     gameModel.GameTitle,
		GameImagePath: gameModel.GameImagePath,
	}

	res := conn.Model(&model.Game{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
