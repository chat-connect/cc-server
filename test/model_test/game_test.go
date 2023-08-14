package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestGameModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		game     *model.Game
		expected bool
	}{
		{
			name:     "Empty Game",
			game:     model.EmptyGame(),
			expected: true,
		},
		{
			name: "Not Empty Game",
			game: &model.Game{
				ID:          1,
				GameKey:     "test_key",
				GenreKey:    "test_key",
				Name:        "test_name",
				Description: "test_explanation",
				Type:        "text",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.game.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
