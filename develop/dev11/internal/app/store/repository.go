package store

import (
	"http-rest-api/internal/app/model"
	"time"
)

// EventRepository ...
type EventRepository interface {
	Create(*model.Event) error
	Update(*model.Event) error
	Delete(id int) error
	FindForDay(date time.Time, user_id int) ([]model.Event, error)
	FindForMonth(month, year string, user_id int) ([]model.Event, error)
	FindForWeek(week, year string, user_id int) ([]model.Event, error)
}
