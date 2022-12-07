package sqlstore_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/sqlstore"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("events")

	s := sqlstore.New(db)
	e := model.TestEvent(t)
	assert.NoError(t, s.Event().Create(model.TestEvent(t)))
	assert.NotNil(t, e)
}

func TestEventRepository_FindByDay(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("events")

	s := sqlstore.New(db)
	var user_id = 1
	var datetime = time.Date(2077, time.January, 1, 0, 0, 0, 0, time.UTC)
	_, err := s.Event().FindForDay(datetime, user_id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	datetime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	e := model.TestEvent(t)
	e.Date = datetime
	s.Event().Create(e)
	e, err = s.Event().FindForDay(datetime, user_id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
}
