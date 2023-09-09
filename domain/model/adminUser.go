package model

import (
    "time"

	"gorm.io/plugin/soft_delete"
)

type AdminUsers []AdminUser

type AdminUser struct {
	ID           int64     `json:"id"`
	AdminUserKey string    `json:"admin_user_key"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Token        string    `json:"token"`
	Status       string    `json:"status"`
	Deleted      soft_delete.DeletedAt `json:"deleted" gorm:"uniqueIndex:udx_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyAdminUser() *AdminUser {
	return &AdminUser{}
}

func (adminUser *AdminUser) IsEmpty() bool {
	return (
		adminUser.ID == 0 &&
		adminUser.AdminUserKey == "" &&
		adminUser.Name == "" &&
		adminUser.Email == "" &&
		adminUser.Password == "" &&
		adminUser.Token == "" &&
		adminUser.Status == "")
}
