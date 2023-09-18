package service

import (
	"log"
	"golang.org/x/crypto/bcrypt"

	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/repository"
)

type UserService interface {
	FindByUserKey(userKey string) (*model.User, error)
	UpdateUser(userKey string, userParam *parameter.UpdateUser) (*model.User, error)
	SearchUser(userKey, name string) (*dto.SearchUsers, error)
}

type userService struct {
	userRepository        repository.UserRepository
	followRepository      repository.FollowRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		userRepository repository.UserRepository,
		followRepository      repository.FollowRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		followRepository:      followRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByUserKey ユーザーを取得する
func (userService *userService) FindByUserKey(userKey string) (*model.User, error) {
	userResult, err := userService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	return userResult, nil
}

// UpdateUser ユーザーを更新する
func (userService *userService) UpdateUser(userKey string, userParam *parameter.UpdateUser) (*model.User, error) {
	checkUser, err := userService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password := string(hashedPassword)

	userModel := &model.User{}
	userModel.UserKey = userKey
	userModel.Email = userParam.Email
	userModel.Name = userParam.Name
	userModel.Password = password
	userModel.Status = checkUser.Status
	userModel.Description = userParam.Description
	userModel.Token = checkUser.Token
	userModel.ImagePath = "/user/" + userKey + ".png"

	// transaction
	tx, err := userService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := userService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := userService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	userResult, err := userService.userRepository.Update(userModel, tx)
	if err != nil {
		return nil, err
	}

	return userResult, nil
}

// SearchUser ユーザーを検索する
func (userService *userService) SearchUser(userKey, name string) (*dto.SearchUsers, error) {
	users, err := userService.userRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	
	userItems := make(dto.SearchUsers, 0, len(*users))
	for _, user := range *users {
		_, err := userService.followRepository.FindByUserKeyAndFollowingUserKey(userKey, user.UserKey)

		result := dto.SearchUser{}
		result.User = user
		if err != nil {
			result.Following = false
		} else {
			result.Following = true
		}

		userItems = append(userItems, result)		
	}

	return &userItems, nil
}
