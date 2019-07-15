package main

import (
	"fmt"
)

type TaskService interface {
	GetList() (tasks []TaskData, err error)
	GetByID(id string) (task TaskData, err error)
	Create(task TaskData) error
	Update(task TaskData) error
	Delete(id string) error
}

type TaskServiceImpl struct {
}

func NewTaskService() TaskService {
	return &TaskServiceImpl{}
}

func (s *TaskServiceImpl) GetList() (tasks []TaskData, err error) {
	fmt.Printf("TaskService.GetList()\n")
	return
}

func (s *TaskServiceImpl) GetByID(id string) (task TaskData, err error) {
	fmt.Printf("TaskService.GetById(%v)\n", id)
	return
}

func (s *TaskServiceImpl) Create(task TaskData) error {
	fmt.Printf("TaskService.Create()\n")
	return nil
}

func (s *TaskServiceImpl) Update(task TaskData) error {
	fmt.Printf("TaskService.Update()\n")
	return nil
}

func (s *TaskServiceImpl) Delete(id string) error {
	fmt.Printf("TaskService.Delete(%v)\n", id)
	return nil
}
