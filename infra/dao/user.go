package dao

import (
	"github.com/chat-connect/cc-server/domain/model"
)

type UserDao struct {
	SqlHandler
}

func (repo *UserDao) FindAll() (users model.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	
	return
}

func (repo *UserDao) FindById(id int) (user model.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}

	return
}

func (repo *UserDao) FindByEmail(email string) (user model.User, err error) {
	if err = repo.Where("email = ?", email).Find(&user).Error; err != nil {
		return
	}

	return
}

func (repo *UserDao) FindByUserKey(userKey string) (user model.User, err error) {
	if err = repo.Where("user_key = ?", userKey).Find(&user).Error; err != nil {
		return
	}

	return
}

func (repo *UserDao) FindByStatus(status string) (users model.Users, err error) {
	if err = repo.Where("status = ?", status).Find(&users).Error; err != nil {
		return
	}

	return
}

func (repo *UserDao) Insert(u model.User) (user model.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	
	user = u

	return
}

func (repo *UserDao) Update(u model.User) (user model.User, err error) {
	if err = repo.Find(&user).Where("user_key = ?", u.UserKey).Update(&u).Error; err != nil {
		return
	}

	return
}

func (repo *UserDao) DeleteByUserKey(user model.User) (err error) {
	if err = repo.Where("user_key = ?", user.UserKey).Delete(&user).Error; err != nil {
		return
	}
	
	return
}
