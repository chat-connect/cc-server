package service

import (
	"github.com/chat-connect/cc-server/domain/repository"
)

type UserService struct {
	UserDao repository.UserRepository
}

func (interactor *UserService) GetOnlineUser() (count int, err error) {
	users, err := interactor.UserDao.FindByStatus("online")
	if err != nil {
		return 0, err
	}

	count = len(users)

	return
}
