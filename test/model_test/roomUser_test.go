package model_test

import (
	"testing"

	"github.com/chat-connect/cc-server/domain/model"
)

func TestRoomUserModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		roomUser *model.RoomUser
		expected bool
	}{
		{
			name:     "Empty RoomUser",
			roomUser: model.EmptyRoomUser(),
			expected: true,
		},
		{
			name:     "Not Empty RoomUser",
			roomUser: &model.RoomUser{
				ID:          1,
				RoomUserKey: "test_key",
				RoomID:      1,
				UserID:      1,
				Host:        false,
				Status:      "online",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.roomUser.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
