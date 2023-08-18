package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type RegisterAdminUser struct {
	AdminUserKey string `json:"admin_user_key"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Message      string `json:"message"`
}

func ToRegisterAdminUser(au *model.AdminUser) *RegisterAdminUser {
	if au == nil {
		return nil
	}

	return &RegisterAdminUser{
		AdminUserKey: au.AdminUserKey,
		Name:         au.Name,
		Email:        au.Email,
		Message:      "user register completed",
	}
}

type EmailValidation struct {
	Message string `json:"message"`
}

func ToEmailValidation() *EmailValidation {
	return &EmailValidation{
		Message: "email already exists",
	}
}

type LoginAdminUser struct {
	AdminUserKey string `json:"admin_user_key"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	Message      string `json:"message"`
}

func ToLoginAdminUser(au *model.AdminUser) *LoginAdminUser {
	return &LoginAdminUser{
		AdminUserKey: au.AdminUserKey,
		Name:         au.Name,
		Email:        au.Email,
		Token:        au.Token,
		Message:      "login admin user completed",
	}
}

type DeleteAdminUser struct {
	Message string `json:"message"`
}

func ToDeleteAdminUser() *DeleteAdminUser {
	return &DeleteAdminUser{
		Message: "delete admin user completed",
	}
}

type CheckAdminUser struct {
	AdminUserKey string `json:"admin_user_key"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Message      string `json:"message"`
}

func ToCheckAdminUser(adminUserKey string, name string, email string) *CheckAdminUser {
	return &CheckAdminUser{
		AdminUserKey: adminUserKey,
		Name:         name,
		Email:        email,
		Message:      "check admin user completed",
	}
}

type LogoutAdminUser struct {
	Message string `json:"message"`
}

func ToLogoutAdminUser() *LogoutAdminUser {
	return &LogoutAdminUser{
		Message: "logout admin user completed",
	}
}
