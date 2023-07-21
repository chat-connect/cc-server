package command

import (
	"fmt"
    "strconv"

	"github.com/chat-connect/cc-server/batch/service"
)

type UserCommand interface {
	GetLoginUser() (err error)
}

type userCommand struct {
	userService service.UserService
}

func NewUserCommand(userService service.UserService) UserCommand {
    return &userCommand{
        userService: userService,
    }
}

// GetLoginUser ログインユーザーを取得する
func (userCommand *userCommand) GetLoginUser() (err error) {
	userResult, err := userCommand.userService.GetLoginUser()
	if err != nil {
		return err
	}

	fmt.Println("Count:" + strconv.FormatInt(userResult, 10))

	return
}
