package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type ChannelChatRepository interface {
	ListByChannelKey(channelKey string) (entity *model.ChannelChats, err error)
	Insert(channelChatModel *model.ChannelChat, tx *gorm.DB) (entity *model.ChannelChat, err error)
	DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error)
}
