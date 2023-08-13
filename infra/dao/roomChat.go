package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type roomChatDao struct {
	Conn *gorm.DB
}

func NewRoomChatDao(conn *gorm.DB) repository.RoomChatRepository {
	return &roomChatDao{
		Conn: conn,
	}
}

func (roomChatDao *roomChatDao) ListByChannelKey(channelKey string) (entity *model.RoomChats, err error) {
	entity = &model.RoomChats{}

	// 最新の100行目までを取得する
	res := roomChatDao.Conn.Where("channel_key = ?", channelKey).Order("created_at DESC").Limit(100).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomChatDao *roomChatDao) Insert(roomChatModel *model.RoomChat, tx *gorm.DB) (entity *model.RoomChat, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomChatDao.Conn
	}

	entity = &model.RoomChat{
		RoomChatKey: roomChatModel.RoomChatKey,
		ChannelKey:  roomChatModel.ChannelKey,
		UserKey:     roomChatModel.UserKey,
		UserName:    roomChatModel.UserName,
		Content:     roomChatModel.Content,
		ImagePath:   roomChatModel.ImagePath,
		PostedAt:    roomChatModel.PostedAt,
	}

	res := conn.Model(&model.RoomChat{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (roomChatDao *roomChatDao) DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomChatDao.Conn
	}
	
	entity := &model.RoomChat{}

	res := conn.Model(&model.RoomChat{}).Where("channel_key IN (?)", channelKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
