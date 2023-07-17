package service

import (
	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/library"
)

type UserService struct {
	UserDao repository.UserRepository
}

func (interactor *UserService) UserById(id int) (user model.User, err error) {
	user, err = interactor.UserDao.FindById(id)

	return
}

func (interactor *UserService) FindByEmail(email string) (user model.User, err error) {
	user, err = interactor.UserDao.FindByEmail(email)

	return
}

func (interactor *UserService) FindByUserKey(userKey string) (user model.User, err error) {
	user, err = interactor.UserDao.FindByUserKey(userKey)

	return
}

func (interactor *UserService) RegisterUser(u model.User) (user model.User, err error) {
	// ユニークキーを生成
	userKey, err := lib.GenerateKey()
	if err != nil {
		return u, err
	}

	u.UserKey = userKey
	u.Status = "offline"

	user, err = interactor.UserDao.Store(u)

	return user, err
}

func (interactor *UserService) UpdateUser(u model.User) (user model.User, err error) {

	user, err = interactor.UserDao.Update(u)

	return user, err
}

func (interactor *UserService) LogoutUser(u model.User) (user model.User, err error) {
	user, err = interactor.UserDao.Update(u)
	
	return user, err
}

func (interactor *UserService) DeleteUser(u model.User) (err error) {
	err = interactor.UserDao.DeleteByUserKey(u)
	
	return
}
