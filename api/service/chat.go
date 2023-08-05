package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/config/key"
)

type ChatService interface {
	ChatCreate(channelKey string, userKey string, chatParam *parameter.ChatCreate) (chatResult *model.Chat, err error)
	ChatList(channelKey string) (chatResult *model.Chats, err error)
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

// ChatList
func (chatService *chatService) ChatList(channelKey string) (chatResult *model.Chats, err error) {
	chatResult, err = chatService.chatRepository.ListByRoomKey(channelKey)
	if err != nil {
		return nil, err
	}

	return chatResult, nil
}

// ChatCreate チャットを作成する
func (chatService *chatService) ChatCreate(channelKey string, userKey string, chatParam *parameter.ChatCreate) (chatResult *model.Chat, err error) {
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

	chatResult, err = chatService.chatRepository.Insert(chatModel, tx)
	if err != nil {
		return nil, err
	}

	return chatResult, nil
}
