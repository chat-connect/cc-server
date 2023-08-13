package service

import (
	"time"
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
	"github.com/game-connect/gc-server/infra/api"
)

type OpenChatService interface {
	ListOpenChat() (openChatResult *model.OpenChats, err error)
	CreateOpenChat(userKey string, openChatParam *parameter.CreateOpenChat) (openChatResult *model.OpenChat, err error)
}

type openChatService struct {
	openChatRepository repository.OpenChatRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewOpenChatService(
		openChatRepository repository.OpenChatRepository,
		userRepository        repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) OpenChatService {
	return &openChatService{
		openChatRepository:    openChatRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// ListOpenChat チャット一覧を取得する
func (openChatService *openChatService) ListOpenChat() (openChatResult *model.OpenChats, err error) {
	openChatResult, err = openChatService.openChatRepository.List()
	if err != nil {
		return nil, err
	}

	return openChatResult, nil
}

// CreateOpenChat チャットを作成する
func (openChatService *openChatService) CreateOpenChat(userKey string, openChatParam *parameter.CreateOpenChat) (openChatResult *model.OpenChat, err error) {
	// transaction
	tx, err := openChatService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := openChatService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := openChatService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	openChatKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	user, err := openChatService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	openChatModel := &model.OpenChat{}
	openChatModel.OpenChatKey = openChatKey
	openChatModel.UserKey = userKey
	openChatModel.UserName = user.Name
	openChatModel.Content = openChatParam.Content
	openChatModel.PostedAt = time.Now()

	if openChatParam.OpenChatImage != nil {
		err = api.UploadImage(*openChatParam.OpenChatImage, openChatKey, "/create_open_chat")
		if err != nil {
			return nil, err
		}

		openChatModel.ImagePath = "/create_open_chat/" + openChatKey + ".png"
	}

	openChatResult, err = openChatService.openChatRepository.Insert(openChatModel, tx)
	if err != nil {
		return nil, err
	}

	return openChatResult, nil
}
