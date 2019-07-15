package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TaskController interface {
	GetList(http.ResponseWriter, *http.Request)
	GetById(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type TaskRouterImpl struct {
	controller TaskController
	m          *mux.Router
}

func NewTaskRouter(controller TaskController) http.Handler {
	r := &TaskRouterImpl{controller, mux.NewRouter()}

	r.m.HandleFunc("/api/tasks", r.controller.GetList).Methods(http.MethodGet)
	r.m.HandleFunc("/api/tasks/{id}", r.controller.GetById).Methods(http.MethodGet)
	r.m.HandleFunc("/api/tasks", r.controller.Create).Methods(http.MethodPost)
	r.m.HandleFunc("/api/tasks/{id}", r.controller.Update).Methods(http.MethodPut)
	r.m.HandleFunc("/api/tasks/{id}", r.controller.Delete).Methods(http.MethodDelete)

	return r.m
}
