package usecase

import (
	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type UserUseCase struct {
	UserDao repository.UserRepository
}

func (interactor *UserUseCase) UserById(id int) (user model.User, err error) {
	user, err = interactor.UserDao.FindById(id)

	return
}

func (interactor *UserUseCase) UserByEmail(email string) (user model.User, err error) {
	user, err = interactor.UserDao.FindByEmail(email)

	return
}

func (interactor *UserUseCase) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserDao.Store(u)

	return
}

func (interactor *UserUseCase) DeleteByUserKey(u model.User) (err error) {
	err = interactor.UserDao.DeleteByUserKey(u)
	
	return
}
