package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/config/key"
)

type RoomUserService interface {
	ListRoomUser(userKey, roomKey string) (*dto.FollowAndUsers, error)
	JoinRoom(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error)
	OutRoom(roomKey string, userKey string) (err error)
}

type roomUserService struct {
	roomRepository        repository.RoomRepository
	roomUserRepository    repository.RoomUserRepository
	userRepository        repository.UserRepository
	followRepository      repository.FollowRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomUserService(
		roomRepository        repository.RoomRepository,
		roomUserRepository    repository.RoomUserRepository,
		userRepository        repository.UserRepository,
		followRepository      repository.FollowRepository,
		transactionRepository repository.TransactionRepository,
	) RoomUserService {
	return &roomUserService{
		roomRepository:        roomRepository,
		roomUserRepository:    roomUserRepository,
		userRepository:        userRepository,
		followRepository:      followRepository,
		transactionRepository: transactionRepository,
	}
}

// ListRoomUser ルームに参加しているユーザー一覧
func (roomUserService *roomUserService) ListRoomUser(userKey, roomKey string) (*dto.FollowAndUsers, error) {
	// ルーム参加中のユーザー一覧
	roomUsers, err := roomUserService.roomUserRepository.ListByRoomKey(roomKey)
	if err != nil {
		return nil, err
	}

	var userKeys []string
	for _, roomUser := range *roomUsers {
		userKeys = append(userKeys, roomUser.UserKey)
	}

	users, err := roomUserService.userRepository.ListByUserKeys(userKeys)
	if err != nil {
		return nil, err
	}

	// フォロー中のユーザー一覧
	follows, err := roomUserService.followRepository.ListByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	var followingUserKeys []string
	for _, follow := range *follows {
		followingUserKeys = append(followingUserKeys, follow.FollowingUserKey)
	}

	followAndUsers := make(dto.FollowAndUsers, 0, len(*roomUsers))
	for _, user := range *users {
		followModel := &model.Follow{}
		for _, follow := range *follows {
			if follow.FollowingUserKey == user.UserKey {
				followModel.FollowKey = follow.FollowKey
				followModel.UserKey = follow.UserKey
				followModel.Mutual = follow.Mutual
			}
		}

		result := dto.FollowAndUser{
			Follow: *followModel,
			User:   user,
		}
		followAndUsers = append(followAndUsers, result)
	}

	return &followAndUsers, nil
}

// JoinRoom ルームに参加する
func (roomUserService *roomUserService) JoinRoom(roomKey string, userKey string) (roomUserResult *model.RoomUser, err error) {
	// transaction
	tx, err := roomUserService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := roomUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	roomUserKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomUserModel := &model.RoomUser{}
	roomUserModel.RoomUserKey = roomUserKey
	roomUserModel.RoomKey = roomKey
	roomUserModel.UserKey = userKey
	roomUserModel.Host = false
	roomUserModel.Status = "online"

	roomUserResult, err = roomUserService.roomUserRepository.Insert(roomUserModel, tx)
	if err != nil {
		return nil, err
	}

	return roomUserResult, nil
}

// OutRoom ルームから退出する
func (roomUserService *roomUserService) OutRoom(roomKey string, userKey string) (err error) {
	// transaction
	tx, err := roomUserService.transactionRepository.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := roomUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	err = roomUserService.roomUserRepository.DeleteByRoomKeyAndUserKey(roomKey, userKey, tx)
	if err != nil {
		return err
	}

	return nil
}
