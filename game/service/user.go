package service

import (	
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/infra/api"
	apiParam "github.com/game-connect/gc-server/infra/api/parameter"
)

type UserService interface {
	LoginUser(userParam *parameter.LoginUser) (*model.User, error)
}

type userService struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		userRepository repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// LoginUser ログイン
func (userService *userService) LoginUser(userParam *parameter.LoginUser) (userResult *model.User, err error) {
	user := &apiParam.LoginUser{}
	user.Email = userParam.Email
	user.Password = userParam.Password

	userResult, err = api.LoginUser(user)
	if err != nil {
		return nil, err
	}

	return userResult, nil
}
