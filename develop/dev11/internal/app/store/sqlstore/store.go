package sqlstore

import (
	"database/sql"
	"http-rest-api/internal/app/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db              *sql.DB
	eventRepository *EventRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Event ...
func (s *Store) Event() store.EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}

	s.eventRepository = &EventRepository{
		store: s,
	}

	return s.eventRepository
}
