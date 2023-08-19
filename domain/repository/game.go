package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameRepository interface {
	FindByGameKey(gameKey string) (entity *model.Game, err error)
	FindByApiKey(apiKey string) (entity *model.Game, err error)
	List() (entity *model.Games, err error)
	Insert(linkgameModel *model.Game, tx *gorm.DB) (entity *model.Game, err error)
}
