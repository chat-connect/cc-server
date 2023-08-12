package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/config/key"
)

type RoomUserService interface {
	JoinRoom(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error)
	OutRoom(roomKey string, userKey string) (err error)
}

type roomUserService struct {
	roomRepository        repository.RoomRepository
	roomUserRepository    repository.RoomUserRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomUserService(
		roomRepository        repository.RoomRepository,
		roomUserRepository    repository.RoomUserRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) RoomUserService {
	return &roomUserService{
		roomRepository:        roomRepository,
		roomUserRepository:    roomUserRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// JoinRoom ルームに参加する
func (roomUserService *roomUserService) JoinRoom(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error) {
	// transaction
	tx, err := roomUserService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := roomUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	roomUserKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomUserModel := &model.RoomUser{}
	roomUserModel.RoomUserKey = roomUserKey
	roomUserModel.RoomKey = roomKey
	roomUserModel.UserKey = userKey
	roomUserModel.Host = false
	roomUserModel.Status = "online"

	roomUserResult, err = roomUserService.roomUserRepository.Insert(roomUserModel, tx)
	if err != nil {
		return nil, err
	}

	return roomUserResult, nil
}

// OutRoom ルームから退出する
func (roomUserService *roomUserService) OutRoom(roomKey string, userKey string) (err error) {
	// transaction
	tx, err := roomUserService.transactionRepository.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := roomUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	err = roomUserService.roomUserRepository.DeleteByRoomKeyAndUserKey(roomKey, userKey, tx)
	if err != nil {
		return err
	}

	return nil
}
