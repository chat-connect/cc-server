package parameter

// room_create
type CreateRoom struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status        string `json:"status"`
}
