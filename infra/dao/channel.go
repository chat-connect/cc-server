package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type channelDao struct {
	Conn *gorm.DB
}

func NewChannelDao(conn *gorm.DB) repository.ChannelRepository {
	return &channelDao{
		Conn: conn,
	}
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
		Explanation: channelModel.Explanation,
		Type:        channelModel.Type,
	}

	res := conn.Model(&model.Channel{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
