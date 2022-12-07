package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

const Layout = "2006-01-02"

type Event struct {
	ID      int       `json:"id"`
	User_id int       `json:"user_id"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
}

// Event ...
func (e *Event) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.Date, validation.Required),
		validation.Field(&e.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&e.User_id, validation.Required),
	)
}
