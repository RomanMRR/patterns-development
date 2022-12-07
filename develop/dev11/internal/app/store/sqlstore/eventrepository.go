package sqlstore

import (
	"database/sql"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"time"
)

// EventRepository ...
type EventRepository struct {
	store *Store
}

// Create ...
func (r *EventRepository) Create(e *model.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO events (user_id, name, date) VALUES ($1, $2, $3) RETURNING id",
		e.User_id,
		e.Name,
		e.Date,
	).Scan(&e.ID)
}

// Update ...
func (r *EventRepository) Update(e *model.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"UPDATE events SET user_id=$1, name=$2, date=$3 WHERE id = $4 RETURNING id",
		e.User_id,
		e.Name,
		e.Date,
		e.ID,
	).Scan(&e.ID)
}

// Update ...
func (r *EventRepository) Delete(id int) error {
	return r.store.db.QueryRow(
		"DELETE FROM events WHERE id = $1 RETURNING id",
		id,
	).Scan(&id)
}

// Find for day ...
func (r *EventRepository) FindForDay(date time.Time, user_id int) ([]model.Event, error) {
	e := &model.Event{}
	rows, err := r.store.db.Query(
		"SELECT id, user_id, name, date FROM events WHERE date = $1 AND user_id = $2",
		date.Format(time.RFC3339),
		user_id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

	}

	var events []model.Event

	for rows.Next() {
		err := rows.Scan(&e.ID, &e.User_id, &e.Name, &e.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, *e)
	}

	return events, nil
}

// Find for month ...
func (r *EventRepository) FindForMonth(month, year string, user_id int) ([]model.Event, error) {
	e := &model.Event{}
	rows, err := r.store.db.Query(
		"SELECT id, user_id, name, date FROM events WHERE extract(month from date) = $1 AND extract(year from date) = $2 and user_id = $3",
		month,
		year,
		user_id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

	}

	var events []model.Event

	for rows.Next() {
		err := rows.Scan(&e.ID, &e.User_id, &e.Name, &e.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, *e)
	}

	return events, nil
}

// Find for week ...
func (r *EventRepository) FindForWeek(week, year string, user_id int) ([]model.Event, error) {
	e := &model.Event{}
	rows, err := r.store.db.Query(
		"SELECT id, user_id, name, date FROM events WHERE extract(week from date) = $1 AND extract(year from date) = $2 and user_id = $3",
		week,
		year,
		user_id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

	}

	var events []model.Event

	for rows.Next() {
		err := rows.Scan(&e.ID, &e.User_id, &e.Name, &e.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, *e)
	}

	return events, nil
}
