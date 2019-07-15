package main

import (
	"encoding/json"
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
	c.r.Methods("GET").Path("/api/tasks/{id}").HandlerFunc(c.getById)
	c.r.Methods("POST").Path("/api/tasks").HandlerFunc(c.create)
	c.r.Methods("PUT").Path("/api/tasks/{id}").HandlerFunc(c.update)
	c.r.Methods("DELETE").Path("/api/tasks/{id}").HandlerFunc(c.delete)

	return c.r
}

func (c *TaskControllerImpl) getList(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.s.GetList()
	writeJSONResponse(w, tasks, err)
}

func (c *TaskControllerImpl) getById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, err := c.s.GetById(id)
	writeJSONResponse(w, task, err)
}

func (c *TaskControllerImpl) create(w http.ResponseWriter, r *http.Request) {
	var task TaskData
	json.NewDecoder(r.Body).Decode(&task)
	c.s.Create(task)
}

func (c *TaskControllerImpl) update(w http.ResponseWriter, r *http.Request) {
	var task TaskData
	json.NewDecoder(r.Body).Decode(&task)
	c.s.Update(task)
}

func (c *TaskControllerImpl) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.s.Delete(id)
}

func writeJSONResponse(w http.ResponseWriter, response interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
