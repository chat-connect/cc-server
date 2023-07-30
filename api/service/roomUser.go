package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/config/key"
)

type RoomUserService interface {
	RoomJoin(roomKey string, userKey string) (*model.RoomUser, error)
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

	roomResult, err := roomUserService.roomRepository.FindByRoomKey(roomKey)
	if err != nil {
		return nil, err
	}

	userResult, err := roomUserService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	roomUserKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomUserModel := &model.RoomUser{}
	roomUserModel.RoomUserKey = roomUserKey
	roomUserModel.RoomID = roomResult.ID
	roomUserModel.UserID = userResult.ID
	roomUserModel.Host = false
	roomUserModel.Status = "online"

	roomUserResult, err = roomUserService.roomUserRepository.Insert(roomUserModel, tx)
	if err != nil {
		return nil, err
	}

	return roomUserResult, nil
}
