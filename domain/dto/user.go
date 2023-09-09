package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type SearchUser struct {
	User      model.User
	Following bool
}

type SearchUsers []SearchUser
