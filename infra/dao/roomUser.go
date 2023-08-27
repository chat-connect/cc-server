package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type roomUserDao struct {
	Conn *gorm.DB
}

func NewRoomUserDao(conn *gorm.DB) repository.RoomUserRepository {
	return &roomUserDao{
		Conn: conn,
	}
}

func (roomUserDao *roomUserDao) FindByRoomKeyAndUserKey(roomKey string, userKey string) (entity *model.RoomUser, err error) {
	entity = &model.RoomUser{}
	
	res := roomUserDao.Conn.Where("room_key = ?", roomKey).Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomUserDao *roomUserDao) ListByUserKey(userKey string) (entity *model.RoomUsers, err error) {
	entity = &model.RoomUsers{}
	
	res := roomUserDao.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomUserDao *roomUserDao) ListByRoomKey(roomKey string) (entity *model.RoomUsers, err error) {
	entity = &model.RoomUsers{}
	
	res := roomUserDao.Conn.Where("room_key = ?", roomKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomUserDao *roomUserDao) Insert(roomUserModel *model.RoomUser, tx *gorm.DB) (entity *model.RoomUser, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomUserDao.Conn
	}

	entity = &model.RoomUser{
		RoomUserKey: roomUserModel.RoomUserKey,
		RoomKey:     roomUserModel.RoomKey,
		UserKey:     roomUserModel.UserKey,
		Host:        roomUserModel.Host,
		Status:      roomUserModel.Status,
	}

	res := conn.Model(&model.RoomUser{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (roomUserDao *roomUserDao) DeleteByRoomKey(roomKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomUserDao.Conn
	}
	
	entity := &model.RoomUser{}

	res := conn.Model(&model.RoomUser{}).Where("room_key = ?", roomKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}

func (roomUserDao *roomUserDao) DeleteByRoomKeyAndUserKey(roomKey string, userKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomUserDao.Conn
	}
	
	entity := &model.RoomUser{}

	res := conn.Model(entity).Where("room_key = ?", roomKey).Where("user_key = ?", userKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
