package interactor

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}
