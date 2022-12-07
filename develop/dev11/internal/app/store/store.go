package store

// Store ...
type Store interface {
	Event() EventRepository
}
