package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type RoomChatRepository interface {
	ListByChannelKey(channelKey string) (entity *model.RoomChats, err error)
	Insert(roomChatModel *model.RoomChat, tx *gorm.DB) (entity *model.RoomChat, err error)
	DeleteByChannelKey(channelKey string, tx *gorm.DB) (err error)
}
