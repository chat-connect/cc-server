package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameScoreRepository interface {
	ListByGameKeyAndUserKey(userKey string, gameKey string, limit int) (*model.GameScores, error)
	Insert(gameScoreModel *model.GameScore, tx *gorm.DB) (entity *model.GameScore, err error)
}
