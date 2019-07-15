package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskControllerImpl struct {
	m *mux.Router
}

func NewTaskController() http.Handler {
	c := &TaskControllerImpl{mux.NewRouter()}

	c.m.Methods("GET").Path("/api/tasks").HandlerFunc(getList)
	c.m.Methods("GET").Path("/api/tasks/{id}").HandlerFunc(getById)
	c.m.Methods("POST").Path("/api/tasks").HandlerFunc(create)
	c.m.Methods("PUT").Path("/api/tasks/{id}").HandlerFunc(update)
	c.m.Methods("DELETE").Path("/api/tasks/{id}").HandlerFunc(delete)

	return c.m
}

func getList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[{"title": "Title 1", "description": "Description 1"}, {"title": "Title 2", "description": "Description 2"}]"`)
}

func getById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"title": "Title 1", "description": "Description 1"}`)
}

func create(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}

func delete(w http.ResponseWriter, r *http.Request) {
}
