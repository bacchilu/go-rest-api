package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{{ID: 1, Name: "Evento1", DateTime: time.Now()}, {ID: 2, Name: "Evento2", DateTime: time.Now()}}

func (e Event) Save() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
