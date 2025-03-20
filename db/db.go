package db

import (
	"database/sql"

	"github.com/bacchilu/rest-api/interactor"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables(DB)

	return DB
}

func createTables(DB *sql.DB) {
	q := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER
		)
	`

	_, err := DB.Exec(q)
	if err != nil {
		panic("Could not create events table.")
	}
}

type sqliteEventRepository struct {
	db *sql.DB
}

func NewSQLiteEventRepository() interactor.DataGateway {
	db := initDB()
	return sqliteEventRepository{db: db}
}

func (r sqliteEventRepository) Create(event interactor.Event) (interactor.Event, error) {
	query := `
		INSERT INTO events (name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return interactor.Event{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return interactor.Event{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return interactor.Event{}, err
	}

	event.ID = id
	return event, nil
}

func (r sqliteEventRepository) GetByID(id int64) (interactor.Event, error) {
	query := "SELECT id, name, description, location, dateTime, userId FROM events WHERE id = ?"
	data := r.db.QueryRow(query, id)
	event := interactor.Event{}
	err := data.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	return event, err
}

func (r sqliteEventRepository) List() ([]interactor.Event, error) {
	query := "SELECT id, name, description, location, dateTime, userId FROM events"
	rows, err := r.db.Query(query)
	if err != nil {
		return []interactor.Event{}, err
	}
	defer rows.Close()

	results := []interactor.Event{}
	for rows.Next() {
		event := interactor.Event{}
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		results = append(results, event)
	}

	return results, nil
}
