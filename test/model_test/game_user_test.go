package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestGameUserModel_EmptyUser(t *testing.T) {
	tests := []struct {
		name     string
		user     *model.GameUser
		expected bool
	}{
		{
			name:     "Empty GameUser",
			user:     model.EmptyGameUser(),
			expected: true,
		},
		{
			name: "Not Empty GameUser",
			user: &model.GameUser{
				ID:          1,
				GameUserKey: "test_game_user_key",
				UserKey:     "test_user_key",
				GameKey:     "test_game_key",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.user.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
