package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
)

type chatDao struct {
	Conn *gorm.DB
}

func NewChatDao(conn *gorm.DB) repository.ChatRepository {
	return &chatDao{
		Conn: conn,
	}
}

func (chatDao *chatDao) Insert(chatModel *model.Chat, tx *gorm.DB) (entity *model.Chat, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = chatDao.Conn
	}

	entity = &model.Chat{
		ChatKey: chatModel.ChatKey,
		RoomKey: chatModel.RoomKey,
		UserKey: chatModel.UserKey,
		Content: chatModel.Content,
	}

	res := conn.Model(&model.Chat{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
