package service

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type UserService interface {
	SearchUser(name string) (*model.Users, error)
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

// SearchUser ユーザーを検索する
func (userService *userService) SearchUser(name string) (*model.Users, error) {
	userResult, err := userService.userRepository.FindByName(name)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}
