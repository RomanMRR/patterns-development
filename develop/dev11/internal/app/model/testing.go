package model

import (
	"testing"
	"time"
)

// Test Event ...
func TestEvent(t *testing.T) *Event {
	var datetime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &Event{
		User_id: 1,
		Name:    "Happy New Year!",
		Date:    datetime,
	}
}
