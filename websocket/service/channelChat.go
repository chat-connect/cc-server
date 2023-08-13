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

type ChannelChatService interface {
	CreateChannelChat(channelKey string, userKey string, channelChatParam *parameter.CreateChannelChat) (chatResult *model.ChannelChat, err error)
}

type channelChatService struct {
	channelChatRepository repository.ChannelChatRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewChannelChatService(
		channelChatRepository repository.ChannelChatRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) ChannelChatService {
	return &channelChatService{
		channelChatRepository: channelChatRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// CreateChannelChat チャットを作成する
func (channelChatService *channelChatService) CreateChannelChat(channelKey string, userKey string, channelChatParam *parameter.CreateChannelChat) (channelChatResult *model.ChannelChat, err error) {
	// transaction
	tx, err := channelChatService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := channelChatService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := channelChatService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	channelChatKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	user, err := channelChatService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	channelChatModel := &model.ChannelChat{}
	channelChatModel.ChannelChatKey = channelChatKey
	channelChatModel.ChannelKey = channelKey
	channelChatModel.UserKey = userKey
	channelChatModel.UserName = user.Name
	channelChatModel.Content = channelChatParam.Content
	channelChatModel.PostedAt = time.Now()

	if channelChatParam.ChannelChatImage != nil {
		err = api.UploadImage(*channelChatParam.ChannelChatImage, channelChatKey, "/create_channel_chat")
		if err != nil {
			return nil, err
		}

		channelChatModel.ImagePath = "/create_channel_chat/" + channelChatKey + ".png"
	}

	channelChatResult, err = channelChatService.channelChatRepository.Insert(channelChatModel, tx)
	if err != nil {
		return nil, err
	}

	return channelChatResult, nil
}
