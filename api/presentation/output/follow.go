package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type ListFollowing struct {
	UserKey string                 `json:"user_key"`
	List    []ListFollowingContent `json:"list"`
	Message string                 `json:"message"`
}

type ListFollowingContent struct {
	FollowKey        string `json:"follow_key"`
	FollowingUserKey string `json:"following_user_key"`
	Mutual           bool   `json:"mutual"`
	MutualFollowKey  string `json:"mutual_follow_key"`
}

func ToListFollowing(userKey string, f *model.Follows) *ListFollowing {
	if f == nil {
		return nil
	}

	var list []ListFollowingContent
	for _, follow := range *f {
		followinContent := ListFollowingContent{
			FollowKey:        follow.FollowKey,
			FollowingUserKey: follow.FollowingUserKey,
			Mutual:           follow.Mutual,
			MutualFollowKey:  follow.MutualFollowKey,
		}
		list = append(list, followinContent)
	}

	return &ListFollowing{
		UserKey: userKey,
		List:    list,
		Message: "list following created",
	}
}

type CreateFollow struct {
	FollowKey        string `json:"follow_key"`
	UserKey          string `json:"user_key"`
	FollowingUserKey string `json:"following_user_key"`
	Mutual           bool   `json:"mutual"`
	MutualFollowKey  string `json:"mutual_follow_key"`
	Message          string `json:"message"`
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
