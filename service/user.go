package service

import (
	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type UserService struct {
	UserDao repository.UserRepository
}

func (interactor *UserService) UserById(id int) (user model.User, err error) {
	user, err = interactor.UserDao.FindById(id)

	return
}

func (interactor *UserService) UserByEmail(email string) (user model.User, err error) {
	user, err = interactor.UserDao.FindByEmail(email)

	return
}

func (interactor *UserService) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserDao.Store(u)

	return
}

func (interactor *UserService) DeleteByUserKey(u model.User) (err error) {
	err = interactor.UserDao.DeleteByUserKey(u)
	
	return
}
