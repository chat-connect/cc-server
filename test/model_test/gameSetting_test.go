package model_test

import (
	"testing"
	"time"

	"github.com/game-connect/gc-server/domain/model"
)

func TestGameSettingModel_EmptySetting(t *testing.T) {
	tests := []struct {
		name       string
		gameSetting *model.GameSetting
		expected   bool
	}{
		{
			name:       "Empty Game Setting",
			gameSetting: model.EmptyGameSetting(),
			expected:   true,
		},
		{
			name: "Not Empty Game Setting",
			gameSetting: &model.GameSetting{
				ID:                 1,
				GameKey:            "test_game_key",
				AdminUserKey:       "test_admin_key",
				GameScore:          true,
				GameComboScore:     true,
				GameRank:           true,
				GamePlayTime:       true,
				GameScoreImagePath: true,
				CreatedAt:          time.Now(),
				UpdatedAt:          time.Now(),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.gameSetting.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}