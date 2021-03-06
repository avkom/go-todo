package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskControllerImpl struct {
	r *mux.Router
	s TaskService
}

func NewTaskController(service TaskService) http.Handler {
	c := &TaskControllerImpl{mux.NewRouter(), service}

	c.r.Methods("GET").Path("/api/tasks").HandlerFunc(c.getList)
	c.r.Methods("GET").Path("/api/tasks/{id}").HandlerFunc(c.getByID)
	c.r.Methods("POST").Path("/api/tasks").HandlerFunc(c.create)
	c.r.Methods("PUT").Path("/api/tasks/{id}").HandlerFunc(c.update)
	c.r.Methods("DELETE").Path("/api/tasks/{id}").HandlerFunc(c.delete)

	return c.r
}

func (c *TaskControllerImpl) getList(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.s.GetList()
	writeJSONResponse(w, tasks, err)
}

func (c *TaskControllerImpl) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, err := c.s.GetByID(id)
	writeJSONResponse(w, task, err)
}

func (c *TaskControllerImpl) create(w http.ResponseWriter, r *http.Request) {
	var task TaskData
	json.NewDecoder(r.Body).Decode(&task)
	err := c.s.Create(task)
	writeEmptyResponse(w, err)
}

func (c *TaskControllerImpl) update(w http.ResponseWriter, r *http.Request) {
	var task TaskData
	json.NewDecoder(r.Body).Decode(&task)
	err := c.s.Update(task)
	writeEmptyResponse(w, err)
}

func (c *TaskControllerImpl) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.s.Delete(id)
	writeEmptyResponse(w, err)
}

func writeJSONResponse(w http.ResponseWriter, response interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func writeEmptyResponse(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
}
