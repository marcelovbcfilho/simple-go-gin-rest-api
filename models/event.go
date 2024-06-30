package models

import (
	"database/sql"
	"go-gin-rest-api/infrastructure"
	"time"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int64     `json:"user_id"`
}

var insertQuery = `INSERT INTO events (name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`
var selectQuery = `SELECT id, name, description, location, datetime, user_id FROM events`

func (event *Event) Save() error {
	stmt, err := infrastructure.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	event.UserId = 1
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}

	event.Id, _ = result.LastInsertId()
	return nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := infrastructure.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	events := make([]Event, 0)

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
