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

func ListGenreKeys(genreKeys *parameter.GenreKeys) (*model.Genres, error) {
	request, err := json.Marshal(genreKeys)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/genre/list_genre_keys", os.Getenv("GC_GAME_URL")), bytes.NewBuffer(request))
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
		Items: &output.ListGenre{},
	}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	genreResults := model.Genres{}
	if listGenre, ok := response.Items.(*output.ListGenre); ok {
		for _, g := range listGenre.List {
			genre := model.Genre{}
			genre.GenreKey = g.GenreKey
			genre.Name = g.Name
			genre.Description = g.Description
			genre.Type = g.Type
			genreResults = append(genreResults, genre)
		}
	}

	return &genreResults, nil
}
