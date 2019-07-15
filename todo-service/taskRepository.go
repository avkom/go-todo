package main

import (
	"fmt"
)

type SQLTaskRepository struct {
}

func NewSQLTaskRepository() TaskRepository {
	return &SQLTaskRepository{}
}

func (s *SQLTaskRepository) GetList() (tasks []TaskData, err error) {
	fmt.Printf("SQLTaskRepository.GetList()\n")
	return
}

func (s *SQLTaskRepository) GetByID(id string) (task TaskData, err error) {
	fmt.Printf("SQLTaskRepository.GetById(%v)\n", id)
	return
}

func (s *SQLTaskRepository) Create(task TaskData) error {
	fmt.Printf("SQLTaskRepository.Create()\n")
	return nil
}

func (s *SQLTaskRepository) Update(task TaskData) error {
	fmt.Printf("SQLTaskRepository.Update()\n")
	return nil
}

func (s *SQLTaskRepository) Delete(id string) error {
	fmt.Printf("SQLTaskRepository.Delete(%v)\n", id)
	return nil
}
