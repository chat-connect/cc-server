package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/infra/api"
	apiParam "github.com/game-connect/gc-server/infra/api/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type UserService interface {
	LoginUser(userParam *parameter.LoginUser) (*model.User, error)
}

type userService struct {
	gameUserRepository    repository.GameUserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		gameUserRepository    repository.GameUserRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		gameUserRepository:    gameUserRepository,
		transactionRepository: transactionRepository,
	}
}

// LoginUser ログイン
func (userService *userService) LoginUser(userParam *parameter.LoginUser) (*model.User, error) {
	// transaction
	tx, err := userService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := userService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := userService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	user := &apiParam.LoginUser{}
	user.Email = userParam.Email
	user.Password = userParam.Password

	userResult, err := api.LoginUser(user)
	if err != nil {
		return nil, err
	}

	gameUserKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	// 初回の場合は連携ユーザーとゲームを記録
	userGame, _ := userService.gameUserRepository.FindByUserKeyAndLinkGameKey(userResult.UserKey, userParam.LinkGameKey)

	if userGame == nil {
		gameUserModel := &model.GameUser{}
		gameUserModel.GameUserKey= gameUserKey
		gameUserModel.UserKey = userResult.UserKey
		gameUserModel.LinkGameKey = userParam.LinkGameKey	
		_, err := userService.gameUserRepository.Insert(gameUserModel, tx)
		if err != nil {
			return nil, err
		}
	}

	return userResult, nil
}
