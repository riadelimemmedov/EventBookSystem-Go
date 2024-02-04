package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// !InitDB
func InitDB() {
	//Connect to sqlite3 and create api.db database
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	//check error if occur
	if err != nil {
		throwError("Could not connect to database")
	}

	//Set max open connection limit size to database
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	//Create table if databaze initialize without problem
	createTable()
}

// !createTable
func createTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events  (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
	`

	_, err := DB.Exec(createEventsTable)

	//Check error if occur when create event table
	if err != nil {
		throwError("Could not create events table")
	}
}

// *throwError
func throwError(message string) {
	//throw panic and close and database terminated and return error to client
	panic(message)
}
