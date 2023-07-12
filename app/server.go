package main

import (
	"context"
	"encoding/json"
	"net/http"
	"test_task/repo"
	"test_task/model"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	SessionName        = "test_restapi"
	CtxKeyUser  ctxKey = iota
	CtxKeyRequestID
)


type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
}

// NewServer return default server
func NewServer() *server {
	srv := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
	}

	srv.routerCfg()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routerCfg() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/recipe", s.handlerRecipeCreation()).Methods("POST")
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(CtxKeyRequestID),
		})
		logger.Infof("Started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"Completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (s *server) handlerRecipeCreation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := repo.GetConnection()
		defer c.Close(context.Background())

		recipe := &model.Recipe{}
		if err := json.NewDecoder(r.Body).Decode(recipe); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := repo.CreateRecipie(c, recipe); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, recipe)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}