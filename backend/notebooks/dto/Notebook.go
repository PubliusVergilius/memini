package dto

type Note struct {
	ID         ID     `json:"id"`
	Body       string `json:"body"`
	UsernameID ID     `json:"username_id"`
}
