package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/infra/api/parameter"
	"github.com/game-connect/gc-server/infra/api/output"
	"github.com/game-connect/gc-server/infra/api/response"
)

func ListGameKeys(gameKeys *parameter.GameKeys) (*model.Games, error) {
	request, err := json.Marshal(gameKeys)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/genre/list_game_keys", os.Getenv("GC_GAME_URL")), bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &response.Success{
		Items: &output.ListGame{},
	}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	gameResults := model.Games{}
	if listGame, ok := response.Items.(*output.ListGame); ok {
		for _, g := range listGame.List {
			game := model.Game{}
			game.GameKey = g.GameKey
			game.GameTitle = g.GameTitle
			gameResults = append(gameResults, game)
		}
	}

	return &gameResults, nil
}
