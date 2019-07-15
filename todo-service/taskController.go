package main

import (
	"fmt"
	"net/http"
)

type TaskControllerImpl struct {
}

func NewTaskController() TaskController {
	return &TaskControllerImpl{}
}

func (c *TaskControllerImpl) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[{"title": "Title 1", "description": "Description 1"}, {"title": "Title 2", "description": "Description 2"}]"`)
}

func (c *TaskControllerImpl) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"title": "Title 1", "description": "Description 1"}`)
}

func (c *TaskControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *TaskControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *TaskControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
