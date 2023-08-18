package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type adminUserDao struct {
	Conn *gorm.DB
}

func NewAdminUserDao(conn *gorm.DB) repository.AdminUserRepository {
	return &adminUserDao{
		Conn: conn,
	}
}

func (adminUserDao *adminUserDao) FindByEmail(email string) (entity *model.AdminUser, err error) {
	entity = &model.AdminUser{}
	res := adminUserDao.Conn.Where("email = ?", email).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (adminUserDao *adminUserDao) FindByAdminUserKey(adminUserKey string) (entity *model.AdminUser, err error) {
	entity = &model.AdminUser{}
	res := adminUserDao.Conn.Where("admin_user_key = ?", adminUserKey).Find(entity)
	if err := res.Error; err != nil {
		return entity, err
	}
	
	return entity, err
}

func (adminUserDao *adminUserDao) Insert(param *model.AdminUser, tx *gorm.DB) (entity *model.AdminUser, err error) {
	entity = &model.AdminUser{
		AdminUserKey: param.AdminUserKey,
		Name:         param.Name,
		Email:        param.Email,
		Password:     param.Password,
		Token:        param.Token,
		Status:       param.Status,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = adminUserDao.Conn
	}
	
	res := conn.Model(&model.AdminUser{}).Create(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (adminUserDao *adminUserDao) Update(param *model.AdminUser, tx *gorm.DB) (entity *model.AdminUser, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = adminUserDao.Conn
	}
	
	entity = &model.AdminUser{
		AdminUserKey: param.AdminUserKey,
		Name:          param.Name,
		Email:         param.Email,
		Password:      param.Password,
		Token:         param.Token,
		Status:        param.Status,
	}

	res := conn.Model(&model.AdminUser{}).Where("admin_user_key = ?", entity.AdminUserKey).Update(entity)
	if err := res.Error; err != nil {
		return entity, err
	}

	return entity, err
}

func (adminUserDao *adminUserDao) DeleteByAdminUserKey(adminUserKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = adminUserDao.Conn
	}
	
	entity := &model.AdminUser{}

	res := conn.Model(&model.User{}).Where("admin_user_key = ?", adminUserKey).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
