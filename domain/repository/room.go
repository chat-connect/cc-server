package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type RoomRepository interface {
	FindByRoomKey(roomKey string) (entity *model.Room, err error)
	ListByRoomKeyList(roomKeyList []string) (entity *model.Rooms, err error)
	Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error)
	DeleteByRoomKey(roomKey string, tx *gorm.DB) (err error)
}
