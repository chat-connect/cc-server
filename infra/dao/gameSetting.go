package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type gameSettingDao struct {
	Conn *gorm.DB
}

func NewGameSettingDao(conn *gorm.DB) repository.GameSettingRepository {
	return &gameSettingDao{
		Conn: conn,
	}
}

func (gameSettingDao *gameSettingDao) FindByGameKey(gameKey string) (*model.GameSetting, error) {
	entity := &model.GameSetting{}
	res := gameSettingDao.Conn.
		Where("game_key = ?", gameKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (gameSettingDao *gameSettingDao) Insert(param *model.GameSetting, tx *gorm.DB) (*model.GameSetting, error) {
	entity := &model.GameSetting{
		GameKey:            param.GameKey,
		AdminUserKey:       param.AdminUserKey,
		GameScore:          param.GameScore,
		GameComboScore:     param.GameComboScore,
		GameRank:           param.GameRank,
		GamePlayTime:       param.GamePlayTime,
		GameScoreImagePath: param.GameScoreImagePath,
	}

	conn := gameSettingDao.Conn
	if tx != nil {
		conn = tx
	}

	res := conn.Model(&model.GameSetting{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, nil
}