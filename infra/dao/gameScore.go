package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type gameScoreDao struct {
	Conn *gorm.DB
}

func NewGameScoreDao(conn *gorm.DB) repository.GameScoreRepository {
	return &gameScoreDao{
		Conn: conn,
	}
}

func (gameScoreDao *gameScoreDao) Insert(param *model.GameScore, tx *gorm.DB) (*model.GameScore, error) {
	entity := &model.GameScore{
		GameScoreKey:       param.GameScoreKey,
		GameKey:            param.GameKey,
		UserKey:            param.UserKey,
		GameUsername:       param.GameUsername,
		GameUserImagePath:  param.GameUserImagePath,
		GameScore:          param.GameScore,
		GameComboScore:     param.GameComboScore,
		GameRank:           param.GameRank,
		GamePlayTime:       param.GamePlayTime,
		GameScoreImagePath: param.GameScoreImagePath,
	}

	conn := gameScoreDao.Conn
	if tx != nil {
		conn = tx
	}

	res := conn.Model(&model.GameScore{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, nil
}