package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestRoomModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		room     *model.Room
		expected bool
	}{
		{
			name:     "Empty Room",
			room:     model.EmptyRoom(),
			expected: true,
		},
		{
			name: "Not Empty Room",
			room: &model.Room{
				ID:          1,
				RoomKey:     "test_key",
				UserKey:     "test_key",
				Name:        "test",
				Description: "test",
				ImagePath:   "/test",
				UserCount:   0,
				Status:      "public",
				Genre:       "RPG",
				Game:        "RPG",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.room.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
