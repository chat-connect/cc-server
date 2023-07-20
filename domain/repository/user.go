package repository

import (
	"github.com/chat-connect/cc-server/domain/model"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindByUserKey(userKey string) (*model.User, error)
	Insert(userModel *model.User) (*model.User, error)
	Update(userModel *model.User) (*model.User, error)
	DeleteByUserKey(userKey string) (error)
}
