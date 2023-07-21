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
	return &userRepository{
		Conn: conn,
	}
}

func (userRepository *userRepository) FindByEmail(email string) (*model.User, error) {
	entity := &model.User{}
	res := userRepository.Conn.Where("email = ?", email).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (userRepository *userRepository) FindByUserKey(userKey string) (*model.User, error) {
	entity := &model.User{}
	res := userRepository.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (userRepository *userRepository) Insert(param *model.User) (*model.User, error) {
	entity := &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	res := userRepository.Conn.Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (userRepository *userRepository) Update(param *model.User) (*model.User, error) {
	entity := &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	res := userRepository.Conn.Model(entity).Where("user_key = ?", entity.UserKey).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (userRepository *userRepository) DeleteByUserKey(userKey string) (error) {
	entity := &model.User{}

	res := userRepository.Conn.Where("user_key = ?", userKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
