package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GameAndGameScore struct {
	Game        model.Game
	GameSetting model.GameSetting
	GameScores  model.GameScores
}

type GameAndGameScores []GameAndGameScore
