package dto

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GameAndGameUser struct {
	Game     model.Game
	GameUser model.GameUser
}

type GameAndGameUsers []GameAndGameUser
