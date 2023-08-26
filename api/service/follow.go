package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type FollowService interface {
	FindByUserKeyAndFollowingUserKey(userKey, followingUserKey string) (*model.Follow, error)
	ListFollowing(userKey string) (*dto.FollowAndUsers, error)
	ListFollowers(userKey string) (*dto.FollowAndUsers, error)
	CreateFollow(userKey string, followParam *parameter.CreateFollow) (*model.Follow, error)
}

type followService struct {
	followRepository      repository.FollowRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewFollowService(
		followRepository      repository.FollowRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) FollowService {
	return &followService{
		followRepository:      followRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByUserKeyAndFollowingUserKey フォローユーザーを取得する
func (followService *followService) FindByUserKeyAndFollowingUserKey(userKey, followingUserKey string) (*model.Follow, error) {
	checkFollow, err := followService.followRepository.FindByUserKeyAndFollowingUserKey(userKey, followingUserKey)
	if err != nil {
		return nil, err
	}

	return checkFollow, nil
}

// ListFollowing フォローしているユーザー一覧
func (followService *followService) ListFollowing(userKey string) (*dto.FollowAndUsers, error) {
	follows, err := followService.followRepository.ListByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	var followingUserKeys []string
	for _, follow := range *follows {
		followingUserKeys = append(followingUserKeys, follow.FollowingUserKey)
	}

	users, err := followService.userRepository.ListByUserKeys(followingUserKeys)
	if err != nil {
		return nil, err
	}

	followAndUsers := make(dto.FollowAndUsers, 0, len(*follows))
	for _, follow := range *follows {
		for _, user := range *users {
			if follow.FollowingUserKey == user.UserKey {
				result := dto.FollowAndUser{
					Follow: follow,
					User:   user,
				}
				followAndUsers = append(followAndUsers, result)
			}
		}
	}

	return &followAndUsers, nil
}

// ListFollowers フォローされているユーザー一覧
func (followService *followService) ListFollowers(userKey string) (*dto.FollowAndUsers, error) {
	follows, err := followService.followRepository.ListByFollowingUserKey(userKey)
	if err != nil {
		return nil, err
	}

	var userKeys []string
	for _, follow := range *follows {
		userKeys = append(userKeys, follow.UserKey)
	}

	users, err := followService.userRepository.ListByUserKeys(userKeys)
	if err != nil {
		return nil, err
	}

	followAndUsers := make(dto.FollowAndUsers, 0, len(*follows))
	for _, follow := range *follows {
		for _, user := range *users {
			if follow.UserKey == user.UserKey {
				result := dto.FollowAndUser{
					Follow: follow,
					User:   user,
				}
				followAndUsers = append(followAndUsers, result)
			}
		}
	}

	return &followAndUsers, nil
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

	// 相互フォローの確認
	checkFollow, _ := followService.followRepository.FindByUserKeyAndFollowingUserKey(followParam.FollowingUserKey, userKey)

	followModel := &model.Follow{}
	followModel.FollowKey = followKey
	followModel.UserKey = userKey
	followModel.FollowingUserKey = followParam.FollowingUserKey
	if checkFollow != nil {
		mutualFollowKey, err := key.GenerateKey()
		if err != nil {
			return nil, err
		}

		followModel.Mutual = true
		followModel.MutualFollowKey = mutualFollowKey
		
		checkFollow.Mutual = true
		checkFollow.MutualFollowKey = mutualFollowKey
		
		_, err = followService.followRepository.Update(checkFollow, tx)
		if err != nil {
			return nil, err
		}
	} else {
		followModel.Mutual = false
		followModel.MutualFollowKey = ""	
	}

	followResult, err := followService.followRepository.Insert(followModel, tx)
	if err != nil {
		return nil, err
	}

	return followResult, nil
}
