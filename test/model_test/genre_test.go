package model_test

import (
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestGenreModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		genre    *model.Genre
		expected bool
	}{
		{
			name:     "Empty Genre",
			genre:  model.EmptyGenre(),
			expected: true,
		},
		{
			name:    "Not Empty Genre",
			genre: &model.Genre{
				ID:          1,
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
			if empty := test.genre.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
