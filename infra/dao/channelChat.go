package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type channelChatDao struct {
	Conn *gorm.DB
}

func NewChannelChatDao(conn *gorm.DB) repository.ChannelChatRepository {
	return &channelChatDao{
		Conn: conn,
	}
}

func (channelChatDao *channelChatDao) ListByChannelKey(channelKey string) (entity *model.ChannelChats, err error) {
	entity = &model.ChannelChats{}

	// 最新の100行目までを取得する
	res := channelChatDao.Conn.Where("channel_key = ?", channelKey).Order("created_at DESC").Limit(100).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (channelChatDao *channelChatDao) Insert(channelChatModel *model.ChannelChat, tx *gorm.DB) (entity *model.ChannelChat, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = channelChatDao.Conn
	}

	entity = &model.ChannelChat{
		ChannelChatKey: channelChatModel.ChannelChatKey,
		ChannelKey:     channelChatModel.ChannelKey,
		UserKey:        channelChatModel.UserKey,
		UserName:       channelChatModel.UserName,
		Content:        channelChatModel.Content,
		ImagePath:      channelChatModel.ImagePath,
		PostedAt:       channelChatModel.PostedAt,
	}

	res := conn.Model(&model.ChannelChat{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (channelChatDao *channelChatDao) DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = channelChatDao.Conn
	}
	
	entity := &model.ChannelChat{}

	res := conn.Model(&model.ChannelChat{}).Where("channel_key IN (?)", channelKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
