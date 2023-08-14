package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type channelDao struct {
	Conn *gorm.DB
}

func NewChannelDao(conn *gorm.DB) repository.ChannelRepository {
	return &channelDao{
		Conn: conn,
	}
}

func (channelDao *channelDao) ListByRoomKey(roomKey string) (entity *model.Channels, err error) {
	entity = &model.Channels{}

	res := channelDao.Conn.Where("room_key = ?", roomKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (channelDao *channelDao) Insert(channelModel *model.Channel, tx *gorm.DB) (entity *model.Channel, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = channelDao.Conn
	}

	entity = &model.Channel{
		ChannelKey:  channelModel.ChannelKey,
		RoomKey:     channelModel.RoomKey,
		Name:        channelModel.Name,
		Description: channelModel.Description,
		Type:        channelModel.Type,
	}

	res := conn.Model(&model.Channel{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (channelDao *channelDao) DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = channelDao.Conn
	}
	
	entity := &model.Channel{}

	res := conn.Model(&model.Channel{}).Where("channel_key IN (?)", channelKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
