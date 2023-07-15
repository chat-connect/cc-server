package repository

import (
    "github.com/chat-connect/cc-server/domain/model"
)

type UserRepository interface {
	FindById(id int) (model.User, error)
	FindByEmail(email string) (model.User, error)
	Store(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	DeleteByUserKey(model.User) error
}
