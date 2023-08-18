package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestLinkGameModel_EmptyLinkGame(t *testing.T) {
	tests := []struct {
		name     string
		linkGame *model.LinkGame
		expected bool
	}{
		{
			name:     "Empty LinkGame",
			linkGame: model.EmptyLinkGame(),
			expected: true,
		},
		{
			name: "Not Empty LinkGame",
			linkGame: &model.LinkGame{
				ID:            1,
				LinkGameKey:   "test_key",
				AdminUserKey:  "test_key",
				ApiKey:        "test_api_key",
				GameTitle:     "test_game_title",
				GameImagePath: "test_image_path",
				GameGenre:     "test_genre",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.linkGame.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}