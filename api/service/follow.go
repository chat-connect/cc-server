package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type FollowService interface {
	CreateFollow(userKey string, followParam *parameter.CreateFollow) (*model.Follow, error)
}

type followService struct {
	followRepository      repository.FollowRepository
	transactionRepository repository.TransactionRepository
}

func NewFollowService(
		followRepository      repository.FollowRepository,
		transactionRepository repository.TransactionRepository,
	) FollowService {
	return &followService{
		followRepository:      followRepository,
		transactionRepository: transactionRepository,
	}
}

// CreateChat チャットを作成する
func (followService *followService) CreateFollow(userKey string, followParam *parameter.CreateFollow) (*model.Follow, error) {
	// transaction
	tx, err := followService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := followService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := followService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	followKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	followModel := &model.Follow{}
	followModel.FollowKey = followKey
	followModel.UserKey = userKey
	followModel.FollowingUserKey = followParam.FollowingUserKey
	followModel.Mutual = false
	followModel.MutualFollowKey = nil

	followResult, err := followService.followRepository.Insert(followModel, tx)
	if err != nil {
		return nil, err
	}

	return followResult, nil
}
