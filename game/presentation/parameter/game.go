package parameter

type CreateGame struct {
	GameTitle string  `json:"game_title"`
	GameImage *string `json:"game_image"`
	GenreKey  string  `json:"genre_key"`
}
