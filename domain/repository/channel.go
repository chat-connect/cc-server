package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
)

type ChannelRepository interface {
	Insert(channelModel *model.Channel, tx *gorm.DB) (entity *model.Channel, err error)
}
