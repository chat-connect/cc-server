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

func (followDao *followDao) CountByUserKey(userKey string) (count int64, err error) {
	entity := &model.Follow{}
	res := followDao.Conn.
		Model(entity).
		Where("user_key = ?", userKey).
		Count(&count)
	if err := res.Error; err != nil {
		return count, err
	}
	
	return count, err
}

func (followDao *followDao) CountByFollowingUserKey(followingUserKey string) (count int64, err error) {
	entity := &model.Follow{}
	res := followDao.Conn.
		Model(entity).
		Where("following_user_key = ?", followingUserKey).
		Count(&count)
	if err := res.Error; err != nil {
		return count, err
	}
	
	return count, err
}

func (followDao *followDao) FindByUserKey(userKey string) (*model.Follow, error) {
	entity := &model.Follow{}
	res := followDao.Conn.
		Where("user_key = ?", userKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (followDao *followDao) FindByUserKeyAndFollowingUserKey(userKey, followingUserKey string) (*model.Follow, error) {
	entity := &model.Follow{}
	res := followDao.Conn.
		Where("user_key = ?", userKey).
		Where("following_user_key = ?", followingUserKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (followDao *followDao) ListByUserKey(userKey string) (*model.Follows, error) {
	entity := &model.Follows{}

	// 最新の100行目までを取得する
	res := followDao.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (followDao *followDao) ListByFollowingUserKey(followingUserKey string) (*model.Follows, error) {
	entity := &model.Follows{}

	// 最新の100行目までを取得する
	res := followDao.Conn.Where("following_user_key = ?", followingUserKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (followDao *followDao) Insert(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = followDao.Conn
	}

	entity := &model.Follow{
		FollowKey:        followModel.FollowKey,
		UserKey:          followModel.UserKey,
		FollowingUserKey: followModel.FollowingUserKey,
		Mutual:           followModel.Mutual,
		MutualFollowKey:  followModel.MutualFollowKey,
	}

	res := conn.Model(&model.Follow{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (followDao *followDao) Update(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = followDao.Conn
	}
	
	entity := &model.Follow{
		FollowKey:        followModel.FollowKey,
		UserKey:          followModel.UserKey,
		FollowingUserKey: followModel.FollowingUserKey,
		Mutual:           followModel.Mutual,
		MutualFollowKey:  followModel.MutualFollowKey,
	}

	res := conn.Model(&model.Follow{}).Where("follow_key = ?", entity.FollowKey).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (followDao *followDao) DeleteByUserKeyAndFollowingUserKey(userKey, followingUserKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = followDao.Conn
	}
	
	entity := &model.Follow{}

	res := conn.Model(&model.Follow{}).Where("user_key = ?", userKey).Where("following_user_key = ?", followingUserKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
