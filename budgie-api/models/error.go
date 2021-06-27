package models

// error model
type Error struct {
	Code    string `json:"code,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}
