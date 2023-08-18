package service

import (
	"log"
	"fmt"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type RoomService interface {
	FindByRoomKey(userKey string) (roomResult *model.Room, err error)
	ListRoom(userKey string) (roomResults *dto.RoomAndGenreAndGames, err error)
	SearchRoom(name string, genre string, game string) (roomResults *dto.RoomAndGenreAndGames, err error)
	CreateRoom(roomParam *parameter.CreateRoom, userKey string) (*model.Room, error)
	DeleteRoom(roomKey string, userKey string) (err error)
}

type roomService struct {
	roomRepository        repository.RoomRepository
	roomUserRepository    repository.RoomUserRepository
	userRepository        repository.UserRepository
	genreRepository       repository.GenreRepository
	gameRepository        repository.GameRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomService(
		roomRepository        repository.RoomRepository,
		roomUserRepository    repository.RoomUserRepository,
		userRepository        repository.UserRepository,
		genreRepository       repository.GenreRepository,
		gameRepository        repository.GameRepository,
		transactionRepository repository.TransactionRepository,
	) RoomService {
	return &roomService{
		roomRepository:        roomRepository,
		roomUserRepository:    roomUserRepository,
		userRepository:        userRepository,
		genreRepository:       genreRepository,
		gameRepository:        gameRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByRoomKey ルームを取得する
func (roomService *roomService) FindByRoomKey(userKey string) (roomResult *model.Room, err error) {
	roomResult, err = roomService.roomRepository.FindByRoomKey(userKey)
	if err != nil {
		return nil, err
	}

	return roomResult, nil
}

// ListRoom ルーム一覧を取得する
func (roomService *roomService) ListRoom(userKey string) (roomResults *dto.RoomAndGenreAndGames, err error) {
	// 参加中のルームを検索
	roomUsers, err := roomService.roomUserRepository.ListByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	var roomKeyList []string
	for _, roomUser := range *roomUsers {
		roomKeyList = append(roomKeyList, roomUser.RoomKey)
	}	

	rooms, err := roomService.roomRepository.ListByRoomKeys(roomKeyList)
	if err != nil {
		return nil, err
	}

	roomItems := make(dto.RoomAndGenreAndGames, 0, len(*rooms))
	for _, room := range *rooms {
		ge, err := roomService.genreRepository.FindByGenreKey(room.Genre)
		if err != nil {
			return nil, err
		}

		ga, err := roomService.gameRepository.FindByGameKey(room.Game)
		if err != nil {
			return nil, err
		}

		result := dto.RoomAndGenreAndGame{
			Room:  room,
			Genre: *ge,
			Game:  *ga,
		}

		roomItems = append(roomItems, result)
	}

	roomResults = &roomItems

	return roomResults, nil
}

// SearchRoom ルームを検索する
func (roomService *roomService) SearchRoom(name string, genre string, game string) (roomResults *dto.RoomAndGenreAndGames, err error) {
	var rooms *model.Rooms

	if name != "" && genre == "" && game == "" {
		// nameのみの場合
		rooms, err = roomService.roomRepository.ListByName(name)
		if err != nil {
			return nil, err
		}
	} else if name == "" && genre != "" && game == "" {
		// genreのみの場合
		rooms, err = roomService.roomRepository.ListByGenre(genre)
		if err != nil {
			return nil, err
		}
	} else if name == "" && genre == "" && game != "" {
		// gameのみの場合
		rooms, err = roomService.roomRepository.ListByGame(game)
		if err != nil {
			return nil, err
		}
	} else if name != "" && genre != "" && game == "" {
		// nameとgenreのみの場合
		rooms, err = roomService.roomRepository.ListByNameAndGenre(name, genre)
		if err != nil {
			return nil, err
		}
	} else if name != "" && genre == "" && game != "" {
		// nameとgameのみの場合
		rooms, err = roomService.roomRepository.ListByNameAndGame(name, game)
		if err != nil {
			return nil, err
		}
	} else if name == "" && genre != "" && game != "" {
		// genreとgameのみの場合
		rooms, err = roomService.roomRepository.ListByGenreAndGame(genre, game)
		if err != nil {
			return nil, err
		}
	} else if name != "" && genre != "" && game != "" {
		// 全ての条件
		rooms, err = roomService.roomRepository.ListByNameAndGenreAndGame(name, genre, game)
		if err != nil {
			return nil, err
		}
	} else { 
		// 何も条件が指定されていない場合
		rooms, err = roomService.roomRepository.List()
		if err != nil {
			return nil, err
		}
	}

	roomItems := make(dto.RoomAndGenreAndGames, 0, len(*rooms))
	for _, room := range *rooms {
		ge, err := roomService.genreRepository.FindByGenreKey(room.Genre)
		if err != nil {
			return nil, err
		}

		ga, err := roomService.gameRepository.FindByGameKey(room.Game)
		if err != nil {
			return nil, err
		}

		result := dto.RoomAndGenreAndGame{
			Room:  room,
			Genre: *ge,
			Game:  *ga,
		}

		roomItems = append(roomItems, result)
	}

	roomResults = &roomItems

	return roomResults, nil
}

// CreateRoom ルームを作成する
func (roomService *roomService) CreateRoom(roomParam *parameter.CreateRoom, userKey string) (roomResult *model.Room, err error) {
	// transaction
	tx, err := roomService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := roomService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	roomKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomModel := &model.Room{}
	roomModel.RoomKey = roomKey
	roomModel.UserKey = userKey
	roomModel.Name = roomParam.Name
	roomModel.Description = roomParam.Description
	roomModel.ImagePath = ""
	roomModel.UserCount = 1
	roomModel.Status = roomParam.Status
	roomModel.Genre = roomParam.Genre
	roomModel.Game = roomParam.Game

	roomResult, err = roomService.roomRepository.Insert(roomModel, tx)
	if err != nil {
		return nil, err
	}

	// ホストユーザーの登録
	roomUserKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomUserModel := &model.RoomUser{}
	roomUserModel.RoomUserKey = roomUserKey
	roomUserModel.RoomKey = roomKey
	roomUserModel.UserKey = userKey
	roomUserModel.Host = true
	roomUserModel.Status = "online"

	_, err = roomService.roomUserRepository.Insert(roomUserModel, tx)
	if err != nil {
		return nil, err
	}

	return roomResult, nil
}

// DeleteRoom ルームを削除する
func (roomService *roomService) DeleteRoom(roomKey string, userKey string) (err error) {
	// transaction
	tx, err := roomService.transactionRepository.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := roomService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	roomUser, err := roomService.roomUserRepository.FindByRoomKeyAndUserKey(roomKey, userKey)
	if err != nil {
		return err
	}

	if !roomUser.Host {
		return fmt.Errorf("user is not host")
	}


	err = roomService.roomUserRepository.DeleteByRoomKey(roomKey, tx)
	if err != nil {
		return err
	}

	err = roomService.roomRepository.DeleteByRoomKey(roomKey, tx)
	if err != nil {
		return err
	}

	return nil
}
