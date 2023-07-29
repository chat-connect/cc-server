package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
)

type RoomUserRepository interface {
	Insert(roomUserModel *model.RoomUser, tx *gorm.DB) (entity *model.RoomUser, err error)
}
