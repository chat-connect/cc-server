package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type RoomRepository interface {
	FindByRoomKey(roomKey string) (entity *model.Room, err error)
	ListByRoomKeys(roomKeys []string) (entity *model.Rooms, err error)
	ListByName(name string) (entity *model.Rooms, err error)
	List() (entity *model.Rooms, err error)
	ListByGenre(genre string) (entity *model.Rooms, err error)
	ListByGame(game string) (entity *model.Rooms, err error)
	ListByNameAndGenre(name string, genre string) (entity *model.Rooms, err error)
	ListByNameAndGame(name string, game string) (entity *model.Rooms, err error)
	ListByGenreAndGame(genre string, game string) (entity *model.Rooms, err error)
	ListByNameAndGenreAndGame(name string, genre string, game string) (entity *model.Rooms, err error)
	Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error)
	DeleteByRoomKey(roomKey string, tx *gorm.DB) (err error)
}
