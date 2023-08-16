package repository

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GameRepository interface {
	List() (entity *model.Games, err error)
}
