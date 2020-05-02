package model

import "time"

// Ticket - ticket model
type Ticket struct {
	Base
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ListID     string     `json:"list_id"`
	Notes      string     `json:"notes"`
	Status     int        `json:"status"`
	Priority   int        `json:"priority"`
	Tags       JSON       `json:"tags"`
	Remind     bool       `json:"remind"`
	RemindDate *time.Time `json:"remind_date"`
	Repeat     int        `json:"repeat"`
}
