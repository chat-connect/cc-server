package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type gameUserDao struct {
	Conn *gorm.DB
}

func NewGameUserDao(conn *gorm.DB) repository.GameUserRepository {
	return &gameUserDao{
		Conn: conn,
	}
}

func (gameUserDao *gameUserDao) FindByUserKeyAndLinkGameKey(userKey string, linkGameKey string) (entity *model.GameUser, err error) {
	entity = &model.GameUser{}
	res := gameUserDao.Conn.Where("user_key = ?", userKey).Where("link_game_key = ?", linkGameKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (gameUserDao *gameUserDao) Insert(param *model.GameUser, tx *gorm.DB) (entity *model.GameUser, err error) {
	entity = &model.GameUser{
		GameUserKey: param.GameUserKey,
		LinkGameKey: param.LinkGameKey,
		UserKey:     param.UserKey,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = gameUserDao.Conn
	}
	
	res := conn.Model(&model.GameUser{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}
