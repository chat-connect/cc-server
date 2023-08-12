package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type ChannelRepository interface {
	ListByRoomKey(roomKey string) (entity *model.Channels, err error)
	Insert(channelModel *model.Channel, tx *gorm.DB) (entity *model.Channel, err error)
	DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error)
}
