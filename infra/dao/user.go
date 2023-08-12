package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type userDao struct {
	Conn *gorm.DB
}

func NewUserDao(conn *gorm.DB) repository.UserRepository {
	return &userDao{
		Conn: conn,
	}
}

func (userDao *userDao) FindByEmail(email string) (entity *model.User, err error) {
	entity = &model.User{}
	res := userDao.Conn.Where("email = ?", email).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (userDao *userDao) FindByUserKey(userKey string) (entity *model.User, err error) {
	entity = &model.User{}
	res := userDao.Conn.Where("user_key = ?", userKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (userDao *userDao) CountByStatus(status string) (count int64, err error) {
	entity := &model.User{}
	res := userDao.Conn.Model(entity).Where("status = ?", status).Count(&count)
	if err := res.Error; err != nil {
		return count, err
	}

	return count, err
}

func (userDao *userDao) Insert(param *model.User, tx *gorm.DB) (entity *model.User, err error) {
	entity = &model.User{
		UserKey:   param.UserKey,
		Name:      param.Name,
		Email:     param.Email,
		Password:  param.Password,
		Token:     param.Token,
		Status:    param.Status,
		ImagePath: param.ImagePath,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userDao.Conn
	}
	
	res := conn.Model(&model.User{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (userDao *userDao) Update(param *model.User, tx *gorm.DB) (entity *model.User, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userDao.Conn
	}
	
	entity = &model.User{
		UserKey:  param.UserKey,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
		Token:    param.Token,
		Status:   param.Status,
	}

	res := conn.Model(&model.User{}).Where("user_key = ?", entity.UserKey).Update(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (userDao *userDao) DeleteByUserKey(userKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = userDao.Conn
	}
	
	entity := &model.User{}

	res := conn.Model(&model.User{}).Where("user_key = ?", userKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
