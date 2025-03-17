package models

import (
	"time"

	"github.com/bacchilu/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{{ID: 1, Name: "Evento1", DateTime: time.Now()}, {ID: 2, Name: "Evento2", DateTime: time.Now()}}

func (e *Event) Save() error {
	id, err := db.InsertEvent(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.GetAllEvents()
	if err != nil {
		return []Event{}, err
	}

	res := []Event{}
	for _, row := range rows {
		res = append(res, Event{ID: row.Id, Name: row.Name, Description: row.Description, Location: row.Location, DateTime: row.DateTime, UserID: row.UserId})
	}

	return res, nil
}
