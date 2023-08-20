package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameSettingRepository interface {
	FindByGameKey(gameKey string) (entity *model.GameSetting, err error)
	Insert(gameSettingModel *model.GameSetting, tx *gorm.DB) (*model.GameSetting, error)
}
