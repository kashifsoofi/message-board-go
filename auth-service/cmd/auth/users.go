package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kashifsoofi/messageboard/auth/pkg/requests"
)

type usersResource struct{}

// Routes creates a REST router for the todos resource
func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware.

	r.Get("/", rs.List)    // GET /users - read a list of users
	r.Post("/", rs.Create) // POST /users - create a new user and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /users/{id} - read a single user by :id
		r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list of stuff.."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		writeResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	request := &requests.CreateUserRequest{}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the request...
	fmt.Fprintf(w, "Request: %+v", request)
	w.Write([]byte("users create"))
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user get"))
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user update"))
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user delete"))
}
