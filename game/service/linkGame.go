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
	CreateLinkGame(adminUserKey string, linkGameParam *parameter.CreateLinkGame) (linkGameResult *model.LinkGame, err error)
}

type linkGameService struct {
	adminUserRepository   repository.AdminUserRepository
	linkGameRepository    repository.LinkGameRepository
	transactionRepository repository.TransactionRepository
}

func NewLinkGameService(
		adminUserRepository   repository.AdminUserRepository,
		linkGameRepository    repository.LinkGameRepository,
		transactionRepository repository.TransactionRepository,
	) LinkGameService {
	return &linkGameService{
		adminUserRepository:   adminUserRepository,
		linkGameRepository:    linkGameRepository,
		transactionRepository: transactionRepository,
	}
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
