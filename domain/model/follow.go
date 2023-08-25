package model

import (
    "time"
)

type Follows []Follow

type Follow struct {
	ID               int64     `json:"id"`
	FollowKey        string    `json:"follow_key"`
	UserKey          string    `json:"user_key"`
	FollowingUserKey string    `json:"following_user_key"`
	Mutual           string    `json:"mutual"`
	MutualFollowKey  string    `json:"mutual_follow_key"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyFollow() *Follow {
	return &Follow{}
}

func (follow *Follow) IsEmpty() bool {
	return (
		follow.ID == 0 &&
		follow.FollowKey == "" &&
		follow.UserKey == "" &&
		follow.FollowingUserKey == "" &&
		follow.Mutual == "" &&
		follow.MutualFollowKey == "")
}
