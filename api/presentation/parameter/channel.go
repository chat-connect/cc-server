package parameter

type CreateChannel struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
