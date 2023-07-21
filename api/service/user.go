package service

import (
	"fmt"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/config/key"
)

type UserService interface {
	FindByEmail(email string) (*model.User, error)
	FindByUserKey(userKey string) (*model.User, error)
	UserRegister(userModel *model.User) (*model.User, error)
	UserLogin(userModel *model.User) (*model.User, error)
	UserCheck(baseToken string) (string, string, string, error)
	UserLogout(userModel *model.User) (*model.User, error)
	UserDelete(userKey string) (error)
}

type userService struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserService(
		userRepository repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) UserService {
	return &userService{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

func (userService *userService) FindByEmail(email string) (userResult *model.User, err error) {
	userResult, err = userService.userRepository.FindByEmail(email)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}

func (userService *userService) FindByUserKey(userKey string) (userResult *model.User, err error) {
	userResult, err = userService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}

func (userService *userService) UserRegister(userModel *model.User) (userResult *model.User, err error) {
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

	userKey, err := key.GenerateKey()
	if err != nil {
		return userModel, err
	}

	userModel.UserKey = userKey
	userModel.Status = "offline"
	userModel.Token = "nil"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	userModel.Password = string(hashedPassword)

	userResult, err = userService.userRepository.Insert(userModel, tx)
	if err != nil {
		return nil, err
	}

	fmt.Println(userResult)

	return userResult, nil
}

func (userService *userService) UserLogin(userModel *model.User) (user *model.User, err error) {
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

	user, err = userService.userRepository.FindByEmail(userModel.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userModel.Password))
	if err != nil {
		return nil, err
	}

	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["user_key"] = user.UserKey
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	token, err := baseToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	user.Token = token
	user.Status = "login"

	_, err = userService.userRepository.Update(user, tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) UserCheck(baseToken string) (userKey string, username string, email string, err error) {
	token, err := jwt.Parse(baseToken[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		
		return []byte("secret"), nil
	})
	if err != nil {
		return userKey, username, email, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userKey, username, email, err
	}

	userKey = claims["user_key"].(string)
	username = claims["username"].(string)
	email = claims["email"].(string)

	return userKey, username, email, nil
}

func (userService *userService) UserLogout(userModel *model.User) (user *model.User, err error) {
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

	user, err = userService.userRepository.Update(userModel, tx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) UserDelete(userKey string) (err error) {
	// transaction
	tx, err := userService.transactionRepository.Begin()
	if err != nil {
		return err
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

	err = userService.userRepository.DeleteByUserKey(userKey, tx)
	if err != nil {
		return err
	}

	return nil
}
