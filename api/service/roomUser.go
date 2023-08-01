package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/config/key"
)

type RoomUserService interface {
	RoomJoin(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error)
	RoomOut(roomKey string, userKey string) (err error)
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

// RoomJoin ルームに参加する
func (roomUserService *roomUserService) RoomJoin(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error) {
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

// RoomOut ルームから退出する
func (roomUserService *roomUserService) RoomOut(roomKey string, userKey string) (err error) {
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
