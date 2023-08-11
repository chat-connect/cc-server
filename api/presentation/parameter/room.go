package parameter

// room_create
type CreateRoom struct {
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
	Status        string `json:"status"`
}
