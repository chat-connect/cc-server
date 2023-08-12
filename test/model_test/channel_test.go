package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestChannelModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		channel  *model.Channel
		expected bool
	}{
		{
			name:     "Empty Channel",
			channel:  model.EmptyChannel(),
			expected: true,
		},
		{
			name:    "Not Empty Channel",
			channel: &model.Channel{
				ID:          1,
				ChannelKey:  "test_key",
				RoomKey:     "test_key",
				Name:        "test_name",
				Explanation: "test_explanation",
				Type:        "text",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.channel.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
