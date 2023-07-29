package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type roomUserRepository struct {
	Conn *gorm.DB
}

func NewRoomUserRepository(conn *gorm.DB) repository.RoomUserRepository {
	return &roomUserRepository{
		Conn: conn,
	}
}

func (roomUserRepository *roomUserRepository) Insert(roomUserModel *model.RoomUser, tx *gorm.DB) (entity *model.RoomUser, err error) {
	entity = &model.RoomUser{
		RoomUserKey: roomUserModel.RoomUserKey,
		RoomID:      roomUserModel.RoomID,
		UserID:      roomUserModel.UserID,
		Host:        roomUserModel.Host,
		Status:      roomUserModel.Status,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomUserRepository.Conn
	}

	res := conn.Model(&model.RoomUser{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
