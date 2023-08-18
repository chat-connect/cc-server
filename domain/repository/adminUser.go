package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type AdminUserRepository interface {
	FindByEmail(email string) (entity *model.AdminUser, err error)
	FindByAdminUserKey(adminUserKey string) (entity *model.AdminUser, err error)
	Insert(adminUserModel *model.AdminUser, tx *gorm.DB) (entity *model.AdminUser, err error)
	Update(adminUserModel *model.AdminUser, tx *gorm.DB) (entity *model.AdminUser, err error)
	DeleteByAdminUserKey(adminUserKey string, tx *gorm.DB) (err error)
}
