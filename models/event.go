package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{{ID: 1, Name: "Evento1", DateTime: time.Now()}, {ID: 2, Name: "Evento2", DateTime: time.Now()}}

func (e Event) Save() {
	e.ID = len(events) + 1
	e.DateTime = time.Now()
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
