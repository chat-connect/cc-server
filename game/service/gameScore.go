package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type GameScoreService interface {
	UpdateGameScore(userKey string, gameScoreParam *parameter.UpdateGameScore) (*model.GameScore, error) 
}

type gameScoreService struct {
	gameScoreRepository   repository.GameScoreRepository
	transactionRepository repository.TransactionRepository
}

func NewGameScoreService(
		gameScoreRepository   repository.GameScoreRepository,
		transactionRepository repository.TransactionRepository,
	) GameScoreService {
	return &gameScoreService{
		gameScoreRepository:   gameScoreRepository,
		transactionRepository: transactionRepository,
	}
}

// CreateLinkGame 連携ゲームを作成する
func (gameScoreService *gameScoreService) UpdateGameScore(userKey string, gameScoreParam *parameter.UpdateGameScore) (*model.GameScore, error) {
	// transaction
	tx, err := gameScoreService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := gameScoreService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := gameScoreService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	gameScoreKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	gameScoreModel := &model.GameScore{}
	gameScoreModel.GameScoreKey = gameScoreKey
	gameScoreModel.LinkGameKey = gameScoreParam.LinkGameKey
	gameScoreModel.UserKey = userKey
	gameScoreModel.GameUsername = gameScoreParam.GameUsername
	gameScoreModel.GameUserImagePath = gameScoreParam.GameUserImage
	gameScoreModel.GameScore = gameScoreParam.GameScore
	gameScoreModel.GameComboScore = gameScoreParam.GameComboScore
	gameScoreModel.GameRank = gameScoreParam.GameRank
	gameScoreModel.GamePlayTime = gameScoreParam.GamePlayTime
	gameScoreModel.GameScoreImagePath = gameScoreParam.GameScoreImage

	gameScoreResult, err := gameScoreService.gameScoreRepository.Insert(gameScoreModel, tx)
	if err != nil {
		return nil, err
	}

	return gameScoreResult, nil
}
