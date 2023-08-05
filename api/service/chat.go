package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/config/key"
)

type ChatService interface {
	ChatCreate(roomKey string, userKey string, chatParam *parameter.ChatCreate) (chatResult *model.Chat, err error)
}

type chatService struct {
	chatRepository        repository.ChatRepository
	transactionRepository repository.TransactionRepository
}

func NewChatService(
		chatRepository        repository.ChatRepository,
		transactionRepository repository.TransactionRepository,
	) ChatService {
	return &chatService{
		chatRepository:        chatRepository,
		transactionRepository: transactionRepository,
	}
}

// ChatCreate チャットを作成する
func (chatService *chatService) ChatCreate(roomKey string, userKey string, chatParam *parameter.ChatCreate) (chatResult *model.Chat, err error) {
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

	chatModel := &model.Chat{}
	chatModel.ChatKey = chatKey
	chatModel.RoomKey = roomKey
	chatModel.UserKey = userKey
	chatModel.Content = chatParam.Content

	chatResult, err = chatService.chatRepository.Insert(chatModel, tx)
	if err != nil {
		return nil, err
	}

	return chatResult, nil
}
