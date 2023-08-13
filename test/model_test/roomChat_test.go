package model_test

import (
	"time"
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestRoomChatModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name        string
		roomChat *model.RoomChat
		expected    bool
	}{
		{
			name:       "Empty Room Chat",
			roomChat: model.EmptyRoomChat(),
			expected:    true,
		},
		{
			name: "Not Empty Room Chat",
			roomChat: &model.RoomChat{
				ID:          1,
				RoomChatKey: "test_key",
				ChannelKey:  "test_key",
				UserKey:     "test_key",
				UserName:    "test_name",
				Content:     "content",
				ImagePath:   "/chat",
				PostedAt:    time.Now(), 
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.roomChat.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
