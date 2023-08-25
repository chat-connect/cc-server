package model_test

import (
	"testing"
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

func TestDirectMailModel_EmptyDirectMail(t *testing.T) {
	tests := []struct {
		name     string
		dm       *model.DirectMail
		expected bool
	}{
		{
			name:     "Empty DirectMail",
			dm:       model.EmptyDirectMail(),
			expected: true,
		},
		{
			name: "Not Empty DirectMail",
			dm: &model.DirectMail{
				ID:              1,
				DirectMailKey:   "test_key",
				MutualFollowKey: "test_follow_key",
				UserKey:         "test_user_key",
				UserName:        "test_user",
				Content:         "test_content",
				ImagePath:       "test_image_path",
				PostedAt:        time.Now(),
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.dm.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
