package service

import (
	"github.com/chat-connect/cc-server/domain/repository"
)

type UserService interface {
	GetLoginUser() (userCount int64, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (userService *userService) GetLoginUser() (userCount int64, err error) {
	userCount, err = userService.userRepository.CountByStatus("login")
	if err != nil {
		return userCount, err
	}

	return userCount , nil
}
