package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
	"github.com/game-connect/gc-server/infra/api"
)

type LinkGameService interface {
	FindByApiKey(apiKey string) (linkGameResult *model.LinkGame, err error)
	CreateLinkGame(adminUserKey string, linkGameParam *parameter.CreateLinkGame) (linkGameResult *model.LinkGame, err error)
}

type linkGameService struct {
	linkGameRepository    repository.LinkGameRepository
	transactionRepository repository.TransactionRepository
}

func NewLinkGameService(
		linkGameRepository    repository.LinkGameRepository,
		transactionRepository repository.TransactionRepository,
	) LinkGameService {
	return &linkGameService{
		linkGameRepository:    linkGameRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByApiKey api_keyからゲームを検索する
func (linkGameService *linkGameService) FindByApiKey(apiKey string) (linkGameResult *model.LinkGame, err error) {
	linkGameResult, err = linkGameService.linkGameRepository.FindByApiKey(apiKey)
	if err != nil {
		return linkGameResult, err
	}

	return linkGameResult, nil
}

// CreateLinkGame 連携ゲームを作成する
func (linkGameService *linkGameService) CreateLinkGame(adminUserKey string, linkGameParam *parameter.CreateLinkGame) (linkGameResult *model.LinkGame, err error) {
	// transaction
	tx, err := linkGameService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := linkGameService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := linkGameService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	linkGameKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	apiKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	linkGameModel := &model.LinkGame{}
	linkGameModel.LinkGameKey = linkGameKey
	linkGameModel.AdminUserKey = adminUserKey
	linkGameModel.ApiKey = apiKey
	linkGameModel.GameTitle = linkGameParam.GameTitle
	linkGameModel.GameImagePath = "/link_game/" + linkGameKey + ".png"
	linkGameModel.GameGenre = linkGameParam.GameGenre

	err = api.UploadImage(*linkGameParam.GameImage, linkGameKey, "/link_game")
	if err != nil {
		return nil, err
	}

	linkGameResult, err = linkGameService.linkGameRepository.Insert(linkGameModel, tx)
	if err != nil {
		return nil, err
	}

	return linkGameResult, nil
}
