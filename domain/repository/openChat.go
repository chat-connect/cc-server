package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type OpenChatRepository interface {
	List() (entity *model.OpenChats, err error)
	Insert(openChatModel *model.OpenChat, tx *gorm.DB) (entity *model.OpenChat, err error)
}
