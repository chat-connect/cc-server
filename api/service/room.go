package service

import (
	"log"
	"fmt"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type RoomService interface {
	FindByRoomKey(userKey string) (roomResult *model.Room, err error)
	ListRoom(userKey string) (roomResult *model.Rooms, err error)
	CreateRoom(roomParam *parameter.CreateRoom, userKey string) (*model.Room, error)
	DeleteRoom(roomKey string, userKey string) (err error)
}

type roomService struct {
	roomRepository        repository.RoomRepository
	roomUserRepository    repository.RoomUserRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomService(
		roomRepository        repository.RoomRepository,
		roomUserRepository    repository.RoomUserRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) RoomService {
	return &roomService{
		roomRepository:        roomRepository,
		roomUserRepository:    roomUserRepository,
		userRepository:        userRepository,
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
func (roomService *roomService) ListRoom(userKey string) (roomResult *model.Rooms, err error) {
	// 参加中のルームを検索
	roomUsers, err := roomService.roomUserRepository.ListByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	var roomKeyList []string
	for _, roomUser := range *roomUsers {
		roomKeyList = append(roomKeyList, roomUser.RoomKey)
	}	

	roomResult, err = roomService.roomRepository.ListByRoomKeyList(roomKeyList)
	if err != nil {
		return nil, err
	}

	return roomResult, nil
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
