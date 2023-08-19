package service

import (
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/repository"
)

type GameUserService interface {
	ListGameUser(userKey string) (*dto.GameAndGameUsers, error)
}

type gameUserService struct {
	gameRepository        repository.GameRepository
	gameUserRepository    repository.GameUserRepository
	transactionRepository repository.TransactionRepository
}

func NewGameUserService(
		gameRepository        repository.GameRepository,
		gameUserRepository    repository.GameUserRepository,
		transactionRepository repository.TransactionRepository,
	) GameUserService {
	return &gameUserService{
		gameRepository:        gameRepository,
		gameUserRepository:    gameUserRepository,
		transactionRepository: transactionRepository,
	}
}

// ListGameUser 連携中ゲーム一覧を取得する
func (gameUserService *gameUserService) ListGameUser(userKey string) (*dto.GameAndGameUsers, error) {
	gameUsers, err := gameUserService.gameUserRepository.ListByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	gameUserResults := make(dto.GameAndGameUsers, 0, len(*gameUsers))
	for _, gameUser := range *gameUsers {
		ga, err := gameUserService.gameRepository.FindByGameKey(gameUser.GameKey)
		if err != nil {
			return nil, err
		}

		result := dto.GameAndGameUser{
			Game:  *ga,
			GameUser: gameUser,
		}

		gameUserResults = append(gameUserResults, result)
	}

	return &gameUserResults, nil
}
