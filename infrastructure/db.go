package infrastructure

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initialize() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	DB = db

	initializeTables()
}

func initializeTables() {
	eventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(eventsTable)
	if err != nil {
		panic(err)
	}

}
