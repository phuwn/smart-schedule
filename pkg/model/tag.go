package model

// Tag - tag model
type Tag struct {
	Base
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
