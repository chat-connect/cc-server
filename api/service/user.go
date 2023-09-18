package service

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/repository"
)

type UserService interface {
	FindByUserKey(userKey string) (*model.User, error)
	SearchUser(userKey, name string) (*dto.SearchUsers, error)
}

type userService struct {
	userRepository        repository.UserRepository
	followRepository      repository.FollowRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		userRepository repository.UserRepository,
		followRepository      repository.FollowRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		followRepository:      followRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByUserKey ユーザーを取得する
func (userService *userService) FindByUserKey(userKey string) (*model.User, error) {
	userResult, err := userService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	return userResult, nil
}

// SearchUser ユーザーを検索する
func (userService *userService) SearchUser(userKey, name string) (*dto.SearchUsers, error) {
	users, err := userService.userRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	
	userItems := make(dto.SearchUsers, 0, len(*users))
	for _, user := range *users {
		_, err := userService.followRepository.FindByUserKeyAndFollowingUserKey(userKey, user.UserKey)

		result := dto.SearchUser{}
		result.User = user
		if err != nil {
			result.Following = false
		} else {
			result.Following = true
		}

		userItems = append(userItems, result)		
	}

	return &userItems, nil
}
