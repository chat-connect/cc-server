package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type RoomChatRepository interface {
	ListByRoomKey(roomKey string) (entity *model.RoomChats, err error)
	Insert(roomChatModel *model.RoomChat, tx *gorm.DB) (entity *model.RoomChat, err error)
}
