package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
)

type RoomRepository interface {
	Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error)
}
