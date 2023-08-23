package parameter

type UpdateGameScore struct {
	GameKey        string `json:"game_key"`
	ApiKey         string `json:"api_key"`
	GameScore      int    `json:"game_score"`
	GameComboScore int    `json:"game_combo_score"`
	GameRank       int    `json:"game_rank"`
	GamePlayTime   int    `json:"game_play_time"`
	GameScoreImage string `json:"game_score_image"`
}
