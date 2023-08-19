package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GameAndGameScore struct {
	Game       model.Game
	GameScores model.GameScores
}

type GameAndGameScores []GameAndGameScore
