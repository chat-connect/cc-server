package model_test

import (
	"time"
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestChannelChatModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name        string
		channelChat *model.ChannelChat
		expected    bool
	}{
		{
			name:       "Empty Channel Chat",
			channelChat: model.EmptyChannelChat(),
			expected:    true,
		},
		{
			name: "Not Empty Chat",
			channelChat: &model.ChannelChat{
				ID:             1,
				ChannelChatKey: "test_key",
				ChannelKey:     "test_key",
				UserKey:        "test_key",
				UserName:       "test_name",
				Content:        "content",
				ImagePath:      "/chat",
				PostedAt:       time.Now(), 
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.channelChat.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
