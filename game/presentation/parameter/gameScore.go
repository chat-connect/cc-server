package parameter

type UpdateGameScore struct {
	GameKey        string `json:"game_key"`
	ApiKey         string `json:"api_key"`
	GameUsername   string `json:"game_username"`
	GameUserImage  string `json:"game_user_image"`
	GameScore      string `json:"game_score"`
	GameComboScore string `json:"game_combo_score"`
	GameRank       string `json:"game_rank"`
	GamePlayTime   int    `json:"game_play_time"`
	GameScoreImage string `json:"game_score_image"`
}