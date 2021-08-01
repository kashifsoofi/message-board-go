package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type server struct{}

func (s *server) ServeHTTP(rs http.ResponseWriter, r *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/users", usersResource{}.Routes())

	http.ListenAndServe(":8080", r)
}
