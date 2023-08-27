package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type RoomUserRepository interface {
	FindByRoomKeyAndUserKey(roomKey string, userKey string) (entity *model.RoomUser, err error)
	ListByUserKey(userKey string) (entity *model.RoomUsers, err error)
	ListByRoomKey(roomKey string) (entity *model.RoomUsers, err error)
	Insert(roomUserModel *model.RoomUser, tx *gorm.DB) (entity *model.RoomUser, err error)
	DeleteByRoomKey(roomKey string, tx *gorm.DB) (err error)
	DeleteByRoomKeyAndUserKey(roomKey string, userKey string, tx *gorm.DB) (err error)
}
