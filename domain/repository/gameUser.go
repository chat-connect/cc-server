package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameUserRepository interface {
	FindByUserKeyAndLinkGameKey(userKey string, linkGameKey string) (*model.GameUser, error)
	ListByUserKey(userKey string)  (*model.GameUsers, error)
	Insert(gameUserModel *model.GameUser, tx *gorm.DB) (*model.GameUser, error)
}
