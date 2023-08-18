package main

import (
	"github.com/game-connect/gc-server/game/presentation/router"
)

// @title Chat Connect
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8004
// @BasePath /
func main() {
	router.Init()
}
