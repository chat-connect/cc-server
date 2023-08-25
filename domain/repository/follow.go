package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type FollowRepository interface {
	FindByUserKeyAndFollowingUserKey(userKey, followingUserKey string) (*model.Follow, error)
	Insert(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error)
	Update(followModel *model.Follow, tx *gorm.DB) (*model.Follow, error)
}
