package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type LinkGameRepository interface {
	FindByApiKey(apiKey string) (entity *model.LinkGame, err error)
	Insert(linkgameModel *model.LinkGame, tx *gorm.DB) (entity *model.LinkGame, err error)
}
