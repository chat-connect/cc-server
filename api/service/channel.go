package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/config/key"
)

type ChannelService interface {
	ChannelCreate(roomKey string, userKey string, channelParam *parameter.ChannelCreate) (channelResult *model.Channel, err error)
}

type channelService struct {
	channelRepository     repository.ChannelRepository
	transactionRepository repository.TransactionRepository
}

func NewChannelService(
		channelRepository     repository.ChannelRepository,
		transactionRepository repository.TransactionRepository,
	) ChannelService {
	return &channelService{
		channelRepository:     channelRepository,
		transactionRepository: transactionRepository,
	}
}

// ChannelCreate チャンネルを作成する
func (channelService *channelService) ChannelCreate(roomKey string, userKey string, channelParam *parameter.ChannelCreate) (channelResult *model.Channel, err error) {
	// transaction
	tx, err := channelService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := channelService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := channelService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	channelKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	channelModel := &model.Channel{}
	channelModel.ChannelKey = channelKey
	channelModel.RoomKey = roomKey
	channelModel.Name = channelParam.Name
	channelModel.Explanation = channelParam.Explanation
	channelModel.Type = channelParam.Type

	channelResult, err = channelService.channelRepository.Insert(channelModel, tx)
	if err != nil {
		return nil, err
	}

	return channelResult, nil
}
