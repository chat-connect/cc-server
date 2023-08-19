package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameScoreRepository interface {
	Insert(gameScoreModel *model.GameScore, tx *gorm.DB) (entity *model.GameScore, err error)
}
