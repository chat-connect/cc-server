package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type CreateFollow struct {
	FollowKey        string  `json:"follow_key"`
	UserKey          string  `json:"user_key"`
	FollowingUserKey string  `json:"following_user_key"`
	Mutual           bool    `json:"mutual"`
	MutualFollowKey  *string `json:"mutual_follow_key"`
	Message          string  `json:"message"`
}

func ToCreateFollow(f *model.Follow) *CreateFollow {
	if f == nil {
		return nil
	}

	return &CreateFollow{
		FollowKey:        f.FollowKey,
		UserKey:          f.UserKey,
		FollowingUserKey: f.FollowingUserKey,
		Mutual:           f.Mutual,
		MutualFollowKey:  f.MutualFollowKey,
		Message:          "room follow completed",
	}
}
