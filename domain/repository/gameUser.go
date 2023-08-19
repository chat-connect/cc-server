package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameUserRepository interface {
	FindByUserKeyAndLinkGameKey(userKey string, linkGameKey string) (entity *model.GameUser, err error)
	Insert(gameUserModel *model.GameUser, tx *gorm.DB) (entity *model.GameUser, err error)
}
