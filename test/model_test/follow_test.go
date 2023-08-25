package model_test

import (
	"testing"
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

func TestFollowModel_EmptyFollow(t *testing.T) {
	tests := []struct {
		name     string
		follow   *model.Follow
		expected bool
	}{
		{
			name:     "Empty Follow",
			follow:   model.EmptyFollow(),
			expected: true,
		},
		{
			name: "Not Empty Follow",
			follow: &model.Follow{
				ID:               1,
				FollowKey:        "test_follow_key",
				UserKey:          "test_user_key",
				FollowingUserKey: "test_following_user_key",
				Mutual:           "test_mutual",
				MutualFollowKey:  "test_mutual_follow_key",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.follow.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
