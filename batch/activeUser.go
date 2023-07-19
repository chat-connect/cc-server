package main

import (
	"fmt"
	"log"

	batchLog "github.com/chat-connect/cc-server/log"
	"github.com/chat-connect/cc-server/batch/controller"
	"github.com/chat-connect/cc-server/batch/response"
	"github.com/chat-connect/cc-server/infra/database"
)

func main() {
	file := batchLog.GeneratBatchLog()
	log.SetOutput(file)

	userController := controller.NewUserController(database.NewSqlHandler())

	count, err := userController.GetOnlineUser()
	if err != nil {
		log.Println(response.NewError(err))
	}

	fmt.Println(count)
}
