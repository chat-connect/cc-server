package parameter

type CreateFollow struct {
	FollowingUserKey string `json:"following_user_key"`
}

type DeleteFollow struct {
	FollowingUserKey string `json:"following_user_key"`
}
