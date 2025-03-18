package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
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

func InsertEvent(name string, description string, location string, dateTime time.Time, userId int) (int64, error) {
	query := `
		INSERT INTO events (name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, description, location, dateTime, userId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

type Row struct {
	Id                          int64
	Name, Description, Location string
	DateTime                    time.Time
	UserId                      int
}

func GetAllEvents() ([]Row, error) {
	query := "SELECT id, name, description, location, dateTime, userId FROM events"
	rows, err := DB.Query(query)
	if err != nil {
		return []Row{}, err
	}
	defer rows.Close()

	results := []Row{}
	for rows.Next() {
		row := Row{}
		rows.Scan(&row.Id, &row.Name, &row.Description, &row.Location, &row.DateTime, &row.UserId)
		results = append(results, row)
	}

	return results, nil
}

func GetSingleEvent(id int64) (Row, error) {
	query := "SELECT id, name, description, location, dateTime, userId FROM events WHERE id = ?"
	data := DB.QueryRow(query, id)
	row := Row{}
	err := data.Scan(&row.Id, &row.Name, &row.Description, &row.Location, &row.DateTime, &row.UserId)

	return row, err
}
