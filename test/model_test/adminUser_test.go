package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestAdminUserModel_EmptyUser(t *testing.T) {
	tests := []struct {
		name      string
		adminUser *model.AdminUser
		expected  bool
	}{
		{
			name:      "Empty Admin User",
			adminUser: model.EmptyAdminUser(),
			expected:  true,
		},
		{
			name: "Not Empty Admin User",
			adminUser: &model.AdminUser{
				ID:           1,
				AdminUserKey: "test_key",
				Name:         "test",
				Email:        "test@example.com",
				Password:     "test_password",
				Token:        "test_token",
				Status:       "login",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.adminUser.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
