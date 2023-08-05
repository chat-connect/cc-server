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

func (chatDao *chatDao) ListByRoomKey(roomKey string) (entity *model.Chats, err error) {
	entity = &model.Chats{}

	res := chatDao.Conn.Where("room_key = ?", roomKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (chatDao *chatDao) Insert(chatModel *model.Chat, tx *gorm.DB) (entity *model.Chat, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = chatDao.Conn
	}

	entity = &model.Chat{
		ChatKey:  chatModel.ChatKey,
		RoomKey:  chatModel.RoomKey,
		UserKey:  chatModel.UserKey,
		UserName: chatModel.UserName,
		Content:  chatModel.Content,
	}

	res := conn.Model(&model.Chat{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
