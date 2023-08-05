package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
)

type ChatRepository interface {
	ListByChannelKey(channelKey string) (entity *model.Chats, err error)
	Insert(chatModel *model.Chat, tx *gorm.DB) (entity *model.Chat, err error)
}
