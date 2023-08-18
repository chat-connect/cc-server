package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
)

type roomDao struct {
	Conn *gorm.DB
}

func NewRoomDao(conn *gorm.DB) repository.RoomRepository {
	return &roomDao{
		Conn: conn,
	}
}

func (roomDao *roomDao) Find() (entity *model.Room, err error) {
	entity = &model.Room{}
	res := roomDao.Conn.Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) FindByRoomKey(roomKey string) (entity *model.Room, err error) {
	entity = &model.Room{}
	res := roomDao.Conn.
		Where("room_key = ?", roomKey).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) List() (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	res := roomDao.Conn.Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByRoomKeys(roomKeys []string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("room_key IN (?)", roomKeys).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByName(name string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("name LIKE ?", "%" + name + "%").
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByGenre(genre string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("genre = ?", genre).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByGame(game string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("game = ?", game).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByNameAndGenre(name string, genre string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("name LIKE ?", "%" + name + "%").
		Where("genre = ?", genre).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByNameAndGame(name string, game string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("name LIKE ?", "%" + name + "%").
		Where("game = ?", game).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByGenreAndGame(genre string, game string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("genre = ?", genre).
		Where("game = ?", game).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) ListByNameAndGenreAndGame(name string, genre string, game string) (entity *model.Rooms, err error) {
	entity = &model.Rooms{}
	
	res := roomDao.Conn.
		Where("name LIKE ?", "%" + name + "%").
		Where("genre = ?", genre).
		Where("game LIKE ?", game).
		Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (roomDao *roomDao) Insert(roomModel *model.Room, tx *gorm.DB) (entity *model.Room, err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomDao.Conn
	}
	
	entity = &model.Room{
		RoomKey:     roomModel.RoomKey,
		UserKey:     roomModel.UserKey,
		Name:        roomModel.Name,
		Description: roomModel.Description,
		ImagePath:   roomModel.ImagePath,
		UserCount:   roomModel.UserCount,
		Status:      roomModel.Status,
		Genre:       roomModel.Genre,
		Game:        roomModel.Game,
	}

	res := conn.
		Model(&model.Room{}).
		Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (roomDao *roomDao) DeleteByRoomKey(roomKey string, tx *gorm.DB) (err error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = roomDao.Conn
	}
	
	entity := &model.Room{}

	res := conn.Model(&model.Room{}).
		Where("room_key = ?", roomKey).
		Delete(entity)
	if err := res.Error; err != nil {
		return err
	}
	
	return err
}
