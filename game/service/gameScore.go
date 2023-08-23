package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type GameScoreService interface {
	ListGameScore(gameKey string, userKey string) (*dto.GameAndGameScore, error) 
	UpdateGameScore(userKey string, gameScoreParam *parameter.UpdateGameScore) (*model.GameScore, error) 
}

type gameScoreService struct {
	gameRepository        repository.GameRepository
	gameSettingRepository repository.GameSettingRepository
	gameScoreRepository   repository.GameScoreRepository
	transactionRepository repository.TransactionRepository
}

func NewGameScoreService(
		gameRepository        repository.GameRepository,
		gameSettingRepository repository.GameSettingRepository,
		gameScoreRepository   repository.GameScoreRepository,
		transactionRepository repository.TransactionRepository,
	) GameScoreService {
	return &gameScoreService{
		gameRepository:        gameRepository,
		gameSettingRepository: gameSettingRepository,
		gameScoreRepository:   gameScoreRepository,
		transactionRepository: transactionRepository,
	}
}

// ListGameScore ゲームスコアを取得する
func (gameScoreService *gameScoreService) ListGameScore(gameKey string, userKey string) (*dto.GameAndGameScore, error) {
	game, err := gameScoreService.gameRepository.FindByGameKey(gameKey)
	if err != nil {
		return nil, err
	}

	gameSetting, err := gameScoreService.gameSettingRepository.FindByGameKey(gameKey)
	if err != nil {
		return nil, err
	}
	
	gameScores, err := gameScoreService.gameScoreRepository.ListByGameKeyAndUserKey(gameKey, userKey, 10)
	if err != nil {
		return nil, err
	}

	sortGameScore := gameScores.Reverse()
	
	gameAndGameScore := &dto.GameAndGameScore{}
	gameAndGameScore.Game = *game
	gameAndGameScore.GameSetting = *gameSetting
	gameAndGameScore.GameScores = sortGameScore
	
	return gameAndGameScore, nil
}

// UpdateGameScore ゲームスコアを追加する
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
	gameScoreModel.GameKey = gameScoreParam.GameKey
	gameScoreModel.UserKey = userKey
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
