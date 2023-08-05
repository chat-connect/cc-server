package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type roomUserDao struct {
	Conn *gorm.DB
}

func NewRoomUserDao(conn *gorm.DB) repository.RoomUserRepository {
	return &roomUserDao{
		Conn: conn,
	}
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
