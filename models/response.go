package models

// Response is the model for the Response
type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
