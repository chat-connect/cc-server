package repository

import (
	"github.com/chat-connect/cc-server/domain/model"
)

type UserRepository interface {
	FindByEmail(email string) (entity *model.User, err error)
	FindByUserKey(userKey string) (entity *model.User, err error)
	Insert(userModel *model.User) (entity *model.User, err error)
	Update(userModel *model.User) (entity *model.User, err error)
	DeleteByUserKey(userKey string) (err error)
}
