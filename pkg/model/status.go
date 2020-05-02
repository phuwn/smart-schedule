package model

// Status - status model
type Status struct {
	Base
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"status"`
}
