package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/config/key"
)

type RoomService interface {
	RoomList(userKey string) (roomResult *model.Rooms, err error)
	RoomCreate(roomParam *parameter.RoomCreate, userKey string) (*model.Room, error)
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

// RoomList ルーム一覧を取得する
func (roomService *roomService) RoomList(userKey string) (roomResult *model.Rooms, err error) {
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

// RoomCreate ルームを作成する
func (roomService *roomService) RoomCreate(roomParam *parameter.RoomCreate, userKey string) (roomResult *model.Room, err error) {
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
	roomModel.Explanation = roomParam.Explanation
	roomModel.ImagePath = ""
	roomModel.UserCount = 1
	roomModel.Status = "public"

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
