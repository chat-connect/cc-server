package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type FollowAndUser struct {
	Follow model.Follow
	User   model.User
}

type FollowAndUsers []FollowAndUser
