package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

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
