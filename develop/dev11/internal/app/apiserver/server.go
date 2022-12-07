package apiserver

import (
	"encoding/json"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s", r.Method)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) configureRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/create_event", s.handleEventsCreate()).Methods("POST")
	s.router.HandleFunc("/update_event", s.handleEventsUpdate()).Methods("POST")
	s.router.HandleFunc("/delete_event", s.handleEventsDelete()).Methods("POST")
	s.router.HandleFunc("/event_for_day", s.handleEventForDay()).Methods("GET")
	s.router.HandleFunc("/event_for_month", s.handleEventForMonth()).Methods("GET")
	s.router.HandleFunc("/event_for_week", s.handleEventForWeek()).Methods("GET")
}

func (s *server) handleEventsCreate() http.HandlerFunc {
	type request struct {
		User_id int
		Name    string
		Date    time.Time
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.Event{
			User_id: req.User_id,
			Date:    req.Date,
			Name:    req.Name,
		}

		if err := s.store.Event().Create(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, e)
	}
}

func (s *server) handleEventsUpdate() http.HandlerFunc {
	type request struct {
		ID      int
		User_id int
		Name    string
		Date    time.Time
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.Event{
			ID:      req.ID,
			User_id: req.User_id,
			Date:    req.Date,
			Name:    req.Name,
		}

		if err := s.store.Event().Update(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, e)
	}
}

func (s *server) handleEventsDelete() http.HandlerFunc {
	type request struct {
		ID int
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var id int = req.ID

		if err := s.store.Event().Delete(id); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "deleted "+strconv.Itoa(id))
	}
}

func (s *server) handleEventForDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		date := r.URL.Query().Get("date")
		dateResult, err := time.Parse(model.Layout, date)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		e, err := s.store.Event().FindForDay(dateResult, user_id)
		if err != nil {
			s.error(w, r, http.StatusServiceUnavailable, err)
			return
		}
		s.respond(w, r, http.StatusOK, e)

	}
}

func (s *server) handleEventForMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		month := r.URL.Query().Get("month")
		year := r.URL.Query().Get("year")

		e, err := s.store.Event().FindForMonth(month, year, user_id)
		if err != nil {
			s.error(w, r, http.StatusServiceUnavailable, err)
			return
		}
		s.respond(w, r, http.StatusOK, e)

	}
}

func (s *server) handleEventForWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		week := r.URL.Query().Get("week")
		year := r.URL.Query().Get("year")

		e, err := s.store.Event().FindForWeek(week, year, user_id)
		if err != nil {
			s.error(w, r, http.StatusServiceUnavailable, err)
			return
		}
		s.respond(w, r, http.StatusOK, e)

	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		var result = make(map[string]interface{})
		result["result"] = data
		json.NewEncoder(w).Encode(result)
	}
}
