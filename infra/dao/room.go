package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type roomRepository struct {
	Conn *gorm.DB
}

func NewRoomRepository(conn *gorm.DB) repository.RoomRepository {
	return &roomRepository{
		Conn: conn,
	}
}

func (roomRepository *roomRepository) Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error) {
	entity = &model.Room{
		RoomKey:     roomModel.RoomKey,
		UserKey:     roomModel.UserKey,
		UserID:      roomModel.UserID,
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
		conn = roomRepository.Conn
	}

	res := conn.Model(&model.Room{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
