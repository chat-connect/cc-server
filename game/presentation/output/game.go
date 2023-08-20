package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type CreateGame struct {
	ApiKey        string `json:"api_key"`
	GameKey       string `json:"game_key"`
	GenreKey      string `json:"genre_key"`
	AdminUserKey  string `json:"admin_user_key"`
	GameTitle     string `json:"game_title"`
	GameImagePath string `json:"game_image_path"`
	Message       string `json:"message"`
}

func ToCreateGame(g *model.Game) *CreateGame {
	if g == nil {
		return nil
	}

	return &CreateGame{
		ApiKey:        g.ApiKey,
		GameKey:       g.GameKey,
		GenreKey:      g.GameKey,
		AdminUserKey:  g.AdminUserKey,
		GameTitle:     g.GameTitle,
		GameImagePath: g.GameImagePath,
		Message:       "create game completed",
	}
}

type ListGenre struct {
	List    []ListGenreContent `json:"list"`
	Message string             `json:"message"`
}

type ListGenreContent struct {
	GenreKey    string `json:"genre_key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToListGenre(c *model.Genres) *ListGenre {
	if c == nil {
		return nil
	}

	var list []ListGenreContent
	for _, genre := range *c {
		genreContent := ListGenreContent{
			GenreKey:    genre.GenreKey,
			Type:        genre.Type,
			Name:        genre.Name,
			Description: genre.Description,
		}
		list = append(list, genreContent)
	}

	return &ListGenre{
		List:    list,
		Message: "list genre created",
	}
}

type ListGame struct {
	List    []ListGameContent `json:"list"`
	Message string             `json:"message"`
}

type ListGameContent struct {
	GameKey   string `json:"game_key"`
	GenreKey  string `json:"genre_key"`
	GameTitle string `json:"game_title"`
}

func ToListGame(c *model.Games) *ListGame {
	if c == nil {
		return nil
	}

	var list []ListGameContent
	for _, genre := range *c {
		genreContent := ListGameContent{
			GameKey:     genre.GameKey,
			GenreKey:    genre.GenreKey,
			GameTitle:   genre.GameTitle,
		}
		list = append(list, genreContent)
	}

	return &ListGame{
		List:    list,
		Message: "game list created",
	}
}

type ListGenreAndGame struct {
	ListGenre []ListGenreContent `json:"list_genre"`
	ListGame  []ListGameContent  `json:"list_game"`
	Message   string             `json:"message"`
}

func ToListGenreAndGame(ge *model.Genres, ga *model.Games) *ListGenreAndGame {
	if ge == nil {
		return nil
	}

	if ga == nil {
		return nil
	}

	var list_genre []ListGenreContent
	for _, genre := range *ge {
		genreContent := ListGenreContent{
			GenreKey:    genre.GenreKey,
			Type:        genre.Type,
			Name:        genre.Name,
			Description: genre.Description,
		}

		list_genre = append(list_genre, genreContent)
	}

	var list_game []ListGameContent
	for _, game := range *ga {
		gameContent := ListGameContent{
			GameKey:     game.GameKey,
			GenreKey:    game.GenreKey,
			GameTitle:   game.GameTitle,
		}
		list_game = append(list_game, gameContent)
	}

	return &ListGenreAndGame{
		ListGenre: list_genre,
		ListGame:  list_game,
		Message:   "list genre and game created",
	}
}
