package service

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type UserService interface {
	FindByUserKey(userKey string) (*model.User, error)
}

type userService struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		userRepository repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByUserKey ユーザーキーからユーザーを検索する
func (userService *userService) FindByUserKey(userKey string) (userResult *model.User, err error) {
	userResult, err = userService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}
