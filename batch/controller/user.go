package controller

import (
	"github.com/chat-connect/cc-server/infra/dao"
	"github.com/chat-connect/cc-server/batch/service"
)

type UserController struct {
	Interactor service.UserService
}

func NewUserController(sqlHandler dao.SqlHandler) *UserController {
	return &UserController{
		Interactor: service.UserService {
				UserDao: &dao.UserDao {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) GetOnlineUser() (count int, err error) {
	count, err = controller.Interactor.GetOnlineUser()

	return
}