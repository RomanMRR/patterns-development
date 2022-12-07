package model_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"testing"
	"time"
)

func TestEvent_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		e       func() *model.Event
		isValid bool
	}{
		{
			name: "valid",
			e: func() *model.Event {
				return model.TestEvent(t)
			},
			isValid: true,
		},
		{
			name: "empty date",
			e: func() *model.Event {
				e := model.TestEvent(t)
				var t time.Time
				e.Date = t
				return e
			},
			isValid: false,
		},
		{
			name: "empty name",
			e: func() *model.Event {
				e := model.TestEvent(t)
				e.Name = ""
				return e
			},
			isValid: false,
		},
		{
			name: "short name",
			e: func() *model.Event {
				e := model.TestEvent(t)
				e.Name = "srf"
				return e
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.e().Validate())
			} else {
				assert.Error(t, tc.e().Validate())
			}
		})
	}
}
