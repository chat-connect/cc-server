package parameter

type CreateLinkGame struct {
	GameTitle string  `json:"game_title"`
	GameImage *string `json:"game_image"`
	GameGenre string  `json:"game_genre"`
}
