package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn: conn}
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	entity := &model.User{}
	res := u.Conn.Where("email = ?", email).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (u *userRepository) FindByUserKey(userKey string) (*model.User, error) {
	entity := &model.User{}
	res := u.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (u *userRepository) Insert(param *model.User) (*model.User, error) {
	entity := &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	res := u.Conn.Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *userRepository) Update(param *model.User) (*model.User, error) {
	entity := &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	res := u.Conn.Model(entity).Where("user_key = ?", entity.UserKey).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *userRepository) DeleteByUserKey(userKey string) (error) {
	entity := &model.User{}

	res := u.Conn.Where("user_key = ?", userKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
