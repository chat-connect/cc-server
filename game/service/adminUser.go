package service

import (
	"fmt"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	
	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
)

type AdminUserService interface {
	FindByEmail(email string) (adminUserResult *model.AdminUser, err error)
	FindByAdminUserKey(adminUserKey string) (adminUserResult *model.AdminUser, err error)
	RegisterAdminUser(adminUserParam *parameter.RegisterAdminUser) (adminUserResult *model.AdminUser, err error)
	LoginAdminUser(adminUserParam *parameter.LoginAdminUser) (adminUser *model.AdminUser, err error)
	CheckAdminUser(baseToken string) (adminUserKey string, name string, email string, err error)
	LogoutAdminUser(adminUserModel *model.AdminUser) (adminUser *model.AdminUser, err error)
	DeleteAdminUser(adminUserKey string) (err error)
}

type adminUserService struct {
	adminUserRepository   repository.AdminUserRepository
	transactionRepository repository.TransactionRepository
}

func NewAdminUserService(
		adminUserRepository   repository.AdminUserRepository,
		transactionRepository repository.TransactionRepository,
	) AdminUserService {
	return &adminUserService{
		adminUserRepository:   adminUserRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByEmail メールアドレスからユーザーを検索する
func (adminUserService *adminUserService) FindByEmail(email string) (adminUserResult *model.AdminUser, err error) {
	adminUserResult, err = adminUserService.adminUserRepository.FindByEmail(email)
	if err != nil {
		return adminUserResult, err
	}

	return adminUserResult, nil
}

// FindByAdminUserKey ユーザーキーからユーザーを検索する
func (adminUserService *adminUserService) FindByAdminUserKey(adminUserKey string) (adminUserResult *model.AdminUser, err error) {
	adminUserResult, err = adminUserService.adminUserRepository.FindByAdminUserKey(adminUserKey)
	if err != nil {
		return adminUserResult, err
	}

	return adminUserResult, nil
}

// RegisterAdminUser ユーザー登録
func (adminUserService *adminUserService) RegisterAdminUser(adminUserParam *parameter.RegisterAdminUser) (adminUserResult *model.AdminUser, err error) {
	// transaction
	tx, err := adminUserService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := adminUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := adminUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	userKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUserParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password := string(hashedPassword)

	adminUserModel := &model.AdminUser{}
	adminUserModel.AdminUserKey = userKey
	adminUserModel.Email = adminUserParam.Email
	adminUserModel.Name = adminUserParam.Name
	adminUserModel.Password = password
	adminUserModel.Status = "not_linked"
	adminUserModel.Token = "nil"

	adminUserResult, err = adminUserService.adminUserRepository.Insert(adminUserModel, tx)
	if err != nil {
		return nil, err
	}

	return adminUserResult, nil
}

// LoginAdminUser ログイン
func (adminUserService *adminUserService) LoginAdminUser(adminUserParam *parameter.LoginAdminUser) (adminUser *model.AdminUser, err error) {
	// transaction
	tx, err := adminUserService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := adminUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := adminUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	adminUser, err = adminUserService.adminUserRepository.FindByEmail(adminUserParam.Email)
	if err != nil {
		return adminUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(adminUserParam.Password))
	if err != nil {
		return nil, err
	}

	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["admin_user_key"] = adminUser.AdminUserKey
	claims["name"] = adminUser.Name
	claims["email"] = adminUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	token, err := baseToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	adminUser.Token = token

	_, err = adminUserService.adminUserRepository.Update(adminUser, tx)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

// CheckAdminUser トークンとユーザーキーからユーザーを確認する
func (adminUserService *adminUserService) CheckAdminUser(baseToken string) (adminUserKey string, name string, email string, err error) {
	token, err := jwt.Parse(baseToken[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		
		return []byte("secret"), nil
	})
	if err != nil {
		return adminUserKey, name, email, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return adminUserKey, name, email, err
	}

	adminUserKey = claims["admin_user_key"].(string)
	name = claims["name"].(string)
	email = claims["email"].(string)

	return adminUserKey, name, email, nil
}

// LogoutAdminUser ログアウト
func (adminUserService *adminUserService) LogoutAdminUser(adminUserModel *model.AdminUser) (adminUser *model.AdminUser, err error) {
	// transaction
	tx, err := adminUserService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := adminUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := adminUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	adminUser, err = adminUserService.adminUserRepository.Update(adminUserModel, tx)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

// DeleteAdminUser ユーザーを削除する
func (adminUserService *adminUserService) DeleteAdminUser(adminUserKey string) (err error){
	// transaction
	tx, err := adminUserService.transactionRepository.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := adminUserService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := adminUserService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	err = adminUserService.adminUserRepository.DeleteByAdminUserKey(adminUserKey, tx)
	if err != nil {
		return err
	}

	return nil
}
