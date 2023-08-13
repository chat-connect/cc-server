package model_test

import (
	"time"
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestOpenChatModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name        string
		openChat *model.OpenChat
		expected    bool
	}{
		{
			name:       "Empty Open Chat",
			openChat: model.EmptyOpenChat(),
			expected:    true,
		},
		{
			name: "Not Open Chat",
			openChat: &model.OpenChat{
				ID:          1,
				OpenChatKey: "test_key",
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
			if empty := test.openChat.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
