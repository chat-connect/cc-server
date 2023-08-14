package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type ChannelService interface {
	ListChannel(roomKey string) (channelResult *model.Channels, err error)
	CreateChannel(roomKey string, userKey string, channelParam *parameter.CreateChannel) (channelResult *model.Channel, err error)
	DeleteChannel(channelKey string) (err error)
}

type channelService struct {
	channelRepository     repository.ChannelRepository
	chatRepository        repository.ChatRepository
	transactionRepository repository.TransactionRepository
}

func NewChannelService(
		channelRepository     repository.ChannelRepository,
		chatRepository        repository.ChatRepository,
		transactionRepository repository.TransactionRepository,
	) ChannelService {
	return &channelService{
		channelRepository:     channelRepository,
		chatRepository:        chatRepository,
		transactionRepository: transactionRepository,
	}
}

// ListChannel チャンネル一覧を取得する
func (channelService *channelService) ListChannel(roomKey string) (channelResult *model.Channels, err error) {
	channelResult, err = channelService.channelRepository.ListByRoomKey(roomKey)
	if err != nil {
		return nil, err
	}

	return channelResult, nil
}

// CreateChannel チャンネルを作成する
func (channelService *channelService) CreateChannel(roomKey string, userKey string, channelParam *parameter.CreateChannel) (channelResult *model.Channel, err error) {
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
	channelModel.Description = channelParam.Description
	channelModel.Type = channelParam.Type

	channelResult, err = channelService.channelRepository.Insert(channelModel, tx)
	if err != nil {
		return nil, err
	}

	return channelResult, nil
}

// DeleteChannel チャンネルを削除する
func (channelService *channelService) DeleteChannel(channelKey string) (err error) {
	// transaction
	tx, err := channelService.transactionRepository.Begin()
	if err != nil {
		return err
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

	err = channelService.channelRepository.DeleteByChannelKey(channelKey, tx)
	if err != nil {
		return err
	}

	err = channelService.chatRepository.DeleteByChannelKey(channelKey, tx)
	if err != nil {
		return err
	}

	return nil
}
