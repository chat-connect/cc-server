package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type linkGameDao struct {
	Conn *gorm.DB
}

func NewLinkGameDao(conn *gorm.DB) repository.LinkGameRepository {
	return &linkGameDao{
		Conn: conn,
	}
}

func (linkGameDao *linkGameDao) FindByApiKey(apiKey string) (entity *model.LinkGame, err error) {
	entity = &model.LinkGame{}
	res := linkGameDao.Conn.Where("api_key = ?", apiKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (linkGameDao *linkGameDao) Insert(linkGameModel *model.LinkGame, tx *gorm.DB) (entity *model.LinkGame, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = linkGameDao.Conn
	}

	entity = &model.LinkGame{
		LinkGameKey: linkGameModel.LinkGameKey,
		AdminUserKey: linkGameModel.AdminUserKey,
		ApiKey: linkGameModel.ApiKey,
		GameTitle: linkGameModel.GameTitle,
		GameImagePath: linkGameModel.GameImagePath,
		GameGenre: linkGameModel.GameGenre,
	}

	res := conn.Model(&model.LinkGame{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}