package service

import (
	"time"
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/websocket/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
	"github.com/game-connect/gc-server/infra/api"
)

type RoomChatService interface {
	CreateRoomChat(channelKey string, userKey string, roomChatParam *parameter.CreateRoomChat) (chatResult *model.RoomChat, err error)
}

type roomChatService struct {
	roomChatRepository    repository.RoomChatRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomChatService(
		roomChatRepository    repository.RoomChatRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) RoomChatService {
	return &roomChatService{
		roomChatRepository:    roomChatRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// CreateRoomChat チャットを作成する
func (roomChatService *roomChatService) CreateRoomChat(channelKey string, userKey string, roomChatParam *parameter.CreateRoomChat) (roomChatResult *model.RoomChat, err error) {
	// transaction
	tx, err := roomChatService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := roomChatService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomChatService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	roomChatKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	user, err := roomChatService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	roomChatModel := &model.RoomChat{}
	roomChatModel.RoomChatKey = roomChatKey
	roomChatModel.ChannelKey = channelKey
	roomChatModel.UserKey = userKey
	roomChatModel.UserName = user.Name
	roomChatModel.Content = roomChatParam.Content
	roomChatModel.PostedAt = time.Now()

	if roomChatParam.RoomChatImage != nil {
		err = api.UploadImage(*roomChatParam.RoomChatImage, roomChatKey, "/create_room_chat")
		if err != nil {
			return nil, err
		}

		roomChatModel.ImagePath = "/create_room_chat/" + roomChatKey + ".png"
	}

	roomChatResult, err = roomChatService.roomChatRepository.Insert(roomChatModel, tx)
	if err != nil {
		return nil, err
	}

	return roomChatResult, nil
}
