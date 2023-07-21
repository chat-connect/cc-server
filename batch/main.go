package main

import (
    "flag"
    "fmt"
    "log"

    "github.com/chat-connect/cc-server/batch/di"
    batchLog "github.com/chat-connect/cc-server/log"
)

func main() {
    // di: wire ./batch/di/wire.go
    userCommand := di.InitializeUserCommand()

    // batch list
    functions := map[string]func() error{
        "GetLoginUser": userCommand.GetLoginUser,
    }

    // receive batch
    commandPtr := flag.String("command", "", "Specify the command")
    flag.Parse()
    command := *commandPtr
    selectedCommand := command

    // run batch
    if function, ok := functions[selectedCommand]; ok {
        logFile := batchLog.GenerateBatchLog()
        log.SetOutput(logFile)

        err := function()
        if err != nil {
            log.Println(err)
        }
    } else {
        fmt.Println("Invalid command")
    }
}