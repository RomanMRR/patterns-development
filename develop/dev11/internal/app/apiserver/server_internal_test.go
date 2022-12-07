package apiserver

import (
	"bytes"
	"encoding/json"
	"http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleEventsCreate(t *testing.T) {
	s := newServer(teststore.New())
	type DataRequest struct {
		User_id int
		Name    string
		Date    time.Time
	}
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: DataRequest{
				User_id: 1,
				Name:    "Some_dataewvew",
				Date:    time.Now(),
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: DataRequest{
				User_id: 1,
				Name:    "Some_dataewvew",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/create_event", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
