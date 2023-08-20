package parameter

type CreateGame struct {
    GameTitle string  `json:"game_title"`
    GameImage *string `json:"game_image"`
    GenreKey  string  `json:"genre_key"`
    Setting   struct {
        GameScore      bool `json:"game_score"`
        GameComboScore bool `json:"game_combo_score"`
        GameRank       bool `json:"game_rank"`
        GamePlayTime   bool `json:"game_play_time"`
        GameScoreImage bool `json:"game_score_image"`
    } `json:"setting"`
}
