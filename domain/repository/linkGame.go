package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type LinkGameRepository interface {
	Insert(linkgameModel *model.LinkGame, tx *gorm.DB) (entity *model.LinkGame, err error)
}
