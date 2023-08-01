package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type roomDao struct {
	Conn *gorm.DB
}

func NewRoomDao(conn *gorm.DB) repository.RoomRepository {
	return &roomDao{
		Conn: conn,
	}
}

func (roomDao *roomDao) FindByRoomKey(roomKey string) (entity *model.Room, err error) {
	entity = &model.Room{}
	res := roomDao.Conn.Where("room_key = ?", roomKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error) {
	entity = &model.Room{
		RoomKey:     roomModel.RoomKey,
		UserKey:     roomModel.UserKey,
		Name:        roomModel.Name,
		Explanation: roomModel.Explanation,
		ImagePath:   roomModel.ImagePath,
		UserCount:   roomModel.UserCount,
		Status:      roomModel.Status,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomDao.Conn
	}

	res := conn.Model(&model.Room{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
