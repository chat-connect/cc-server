package service

import (
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type GameService interface {
	ListGame() (gameResult *model.Games, err error)
}

type gameService struct {
	gameRepository repository.GameRepository
}

func NewGameService(
		gameRepository repository.GameRepository,
	) GameService {
	return &gameService{
		gameRepository: gameRepository,
	}
}

// ListGame ゲーム一覧を取得する
func (gameService *gameService) ListGame() (genreResult *model.Games, err error) {
	genreResult, err = gameService.gameRepository.List()
	if err != nil {
		return nil, err
	}

	return genreResult, nil
}
