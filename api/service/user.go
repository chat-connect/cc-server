package service

import (
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/lib"
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
	userRepo repository.UserRepository
}

func NewTaskService(userRepo repository.UserRepository) UserService {
	return &userService{ userRepo: userRepo }
}

func (u *userService) FindByEmail(email string) (*model.User, error) {
	userResult, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}

func (u *userService) FindByUserKey(userKey string) (*model.User, error) {
	userResult, err := u.userRepo.FindByUserKey(userKey)
	if err != nil {
		return userResult, err
	}

	return userResult, nil
}

func (u *userService) UserRegister(userModel *model.User) (*model.User, error) {
	userKey, err := lib.GenerateKey()
	if err != nil {
		return userModel, err
	}

	userModel.UserKey = userKey
	userModel.Status = "offline"
	userModel.Token = "nil"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	userModel.Password = string(hashedPassword)

	userResult, err := u.userRepo.Insert(userModel)
	if err != nil {
		return nil, err
	}

	return userResult, nil
}

func (u *userService) UserLogin(userModel *model.User) (*model.User, error) {
	user, err := u.userRepo.FindByEmail(userModel.Email)
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
	user.Status = "online"

	_, err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UserCheck(baseToken string) (userKey string, username string, email string, err error) {
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

func (u *userService) UserLogout(userModel *model.User) (*model.User, error) {
	user, err := u.userRepo.Update(userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UserDelete(userKey string) (error) {
	err := u.userRepo.DeleteByUserKey(userKey)
	if err != nil {
		return err
	}

	return nil
}
