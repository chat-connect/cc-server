package model

import (
	"testing"
)

func TestUserModel_EmptyUser(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		expected bool
	}{
		{
			name:     "Empty User",
			user:     EmptyUser(),
			expected: true,
		},
		{
			name: "Not Empty User",
			user: &User{
				ID:        1,
				UserKey:   "test_key",
				Username:  "test",
				Email:     "test@example.com",
				Password:  "test_password",
				Token:     "test_token",
				Status:    "login",
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
