package models

import (
	"context"
	"github.com/erik-olsson-op/go-rest/internal/database"
	"time"
)

type Event struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() (int64, error) {
	query := "INSERT INTO event (title,description,location,date_time,user_id) VALUES (?,?,?,?,?)"
	prepare, err := database.Connection.PrepareContext(context.Background(), query)
	if err != nil {
		return -1, err
	}
	defer prepare.Close()
	result, err := prepare.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM event"
	rows, err := database.Connection.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id string) (*Event, error) {
	query := "SELECT * FROM event WHERE id = ?"
	row := database.Connection.QueryRow(query, id)
	var event Event
	err := row.Scan(
		&event.Id,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func DeleteEventById(id string) error {
	query := "DELETE FROM event WHERE id = ?"
	prepare, err := database.Connection.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.ExecContext(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEventById(id string, e Event) error {
	query := `UPDATE event
              SET title = ?, description = ?, location = ?, date_time = ?
              WHERE id = ?`
	prepare, err := database.Connection.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.ExecContext(context.Background(), e.Title, e.Description, e.Location, e.DateTime, id)
	if err != nil {
		return err
	}
	return nil
}
