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

func (userRepository *userRepository) FindByEmail(email string) (entity *model.User, err error) {
	entity = &model.User{}
	res := userRepository.Conn.Where("email = ?", email).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (userRepository *userRepository) FindByUserKey(userKey string) (entity *model.User, err error) {
	entity = &model.User{}
	res := userRepository.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, nil
}

func (userRepository *userRepository) CountByStatus(status string) (count int64, err error) {
	entity := &model.User{}
	res := userRepository.Conn.Model(entity).Where("status = ?", status).Count(&count)
	if err := res.Error; err != nil {
		return count, err
	}

	return count, nil
}

func (userRepository *userRepository) Insert(param *model.User, tx *gorm.DB) (entity *model.User, err error) {
	entity = &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userRepository.Conn
	}
	
	res := conn.Model(&model.User{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (userRepository *userRepository) Update(param *model.User, tx *gorm.DB) (entity *model.User, err error) {
	entity = &model.User{
		UserKey:  param.UserKey,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userRepository.Conn
	}
	
	res := conn.Model(&model.User{}).Where("user_key = ?", entity.UserKey).Update(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (userRepository *userRepository) DeleteByUserKey(userKey string, tx *gorm.DB) (err error) {
	entity := &model.User{}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userRepository.Conn
	}
	
	res := conn.Model(&model.User{}).Where("user_key = ?", userKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
