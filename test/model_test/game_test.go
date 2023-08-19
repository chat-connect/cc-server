package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestLinkGameModel_EmptyLinkGame(t *testing.T) {
	tests := []struct {
		name     string
		game *model.Game
		expected bool
	}{
		{
			name:     "Empty Game",
			game: model.EmptyGame(),
			expected: true,
		},
		{
			name: "Not Empty Game",
			game: &model.Game{
				ID:            1,
				GameKey:       "test_key",
				GenreKey:      "test_genre",
				AdminUserKey:  "test_key",
				ApiKey:        "test_api_key",
				GameTitle:     "test_game_title",
				GameImagePath: "test_image_path",
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