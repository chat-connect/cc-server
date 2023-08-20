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

func (gameScoreDao *gameScoreDao) ListByGameKeyAndUserKey(userKey string, gameKey string) (*model.GameScores, error) {
	entity := &model.GameScores{}
	res := gameScoreDao.Conn.
		Where("game_key = ?", userKey).
		Where("user_key = ?", gameKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (gameScoreDao *gameScoreDao) Insert(param *model.GameScore, tx *gorm.DB) (*model.GameScore, error) {
	entity := &model.GameScore{
		GameScoreKey:       param.GameScoreKey,
		GameKey:            param.GameKey,
		UserKey:            param.UserKey,
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