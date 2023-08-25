package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type FollowRepository interface {
	Insert(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error)
}
