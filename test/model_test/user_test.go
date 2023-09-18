package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestUserModel_EmptyUser(t *testing.T) {
	tests := []struct {
		name     string
		user     *model.User
		expected bool
	}{
		{
			name:     "Empty User",
			user:     model.EmptyUser(),
			expected: true,
		},
		{
			name: "Not Empty User",
			user: &model.User{
				ID:          1,
				UserKey:     "test_key",
				Name:        "test",
				Email:       "test@example.com",
				Password:    "test_password",
				Token:       "test_token",
				Status:      "login",
				Description: "Description",
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
