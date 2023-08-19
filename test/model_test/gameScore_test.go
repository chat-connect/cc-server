package model_test

import (
	"testing"
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

func TestGameScoreModel_EmptyScore(t *testing.T) {
	tests := []struct {
		name     string
		gameScore *model.GameScore
		expected bool
	}{
		{
			name:     "Empty GameScore",
			gameScore: model.EmptyGameScore(),
			expected: true,
		},
		{
			name: "Not Empty GameScore",
			gameScore: &model.GameScore{
				ID:                 1,
				GameScoreKey:       "test_score_key",
				LinkGameKey:        "test_link_game_key",
				GameUsername:       "test_username",
				GameUserImagePath:  "test_image_path",
				GameScore:          "100",
				GameComboScore:     "50",
				GameRank:           "A",
				GamePlayTime:       "10:00",
				GameScoreImagePath: "test_score_image_path",
				CreatedAt:          time.Now(),
				UpdatedAt:          time.Now(),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.gameScore.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
