package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type openChatDao struct {
	Conn *gorm.DB
}

func NewOpenChatDao(conn *gorm.DB) repository.OpenChatRepository {
	return &openChatDao{
		Conn: conn,
	}
}

func (openChatDao *openChatDao) List() (entity *model.OpenChats, err error) {
	entity = &model.OpenChats{}

	// 最新の100行目までを取得する
	res := openChatDao.Conn.Order("created_at DESC").Limit(100).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (openChatDao *openChatDao) Insert(openChatModel *model.OpenChat, tx *gorm.DB) (entity *model.OpenChat, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = openChatDao.Conn
	}

	entity = &model.OpenChat{
		OpenChatKey: openChatModel.OpenChatKey,
		UserKey:     openChatModel.UserKey,
		UserName:    openChatModel.UserName,
		Content:     openChatModel.Content,
		ImagePath:   openChatModel.ImagePath,
		PostedAt:    openChatModel.PostedAt,
	}

	res := conn.Model(&model.OpenChat{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
