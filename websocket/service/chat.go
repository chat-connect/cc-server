package service

import (
	"time"
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/websocket/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type ChatService interface {
	CreateChat(channelKey string, userKey string, chatParam *parameter.CreateChat) (chatResult *model.Chat, err error)
}

type chatService struct {
	chatRepository        repository.ChatRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewChatService(
		chatRepository        repository.ChatRepository,
		userRepository repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) ChatService {
	return &chatService{
		chatRepository:        chatRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// CreateChat チャットを作成する
func (chatService *chatService) CreateChat(channelKey string, userKey string, chatParam *parameter.CreateChat) (chatResult *model.Chat, err error) {
	// transaction
	tx, err := chatService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := chatService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := chatService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	chatKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	user, err := chatService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	chatModel := &model.Chat{}
	chatModel.ChatKey = chatKey
	chatModel.ChannelKey = channelKey
	chatModel.UserKey = userKey
	chatModel.UserName = user.Name
	chatModel.Content = chatParam.Content
	chatModel.PostedAt = time.Now()

	chatResult, err = chatService.chatRepository.Insert(chatModel, tx)
	if err != nil {
		return nil, err
	}

	return chatResult, nil
}
