package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type followDao struct {
	Conn *gorm.DB
}

func NewFollowDao(conn *gorm.DB) repository.FollowRepository {
	return &followDao{
		Conn: conn,
	}
}

func (followDao *followDao) Insert(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = followDao.Conn
	}

	entity := &model.Follow{
		FollowKey:  followModel.FollowKey,
		UserKey:    followModel.UserKey,
		FollowingUserKey:   followModel.FollowingUserKey,
		Mutual:  followModel.Mutual,
		MutualFollowKey:  followModel.MutualFollowKey,
	}

	res := conn.Model(&model.Follow{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
