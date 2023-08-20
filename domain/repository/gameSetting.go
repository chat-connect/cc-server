package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type GameSettingRepository interface {
	Insert(gameSettingModel *model.GameSetting, tx *gorm.DB) (*model.GameSetting, error)
}
