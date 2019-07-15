package main

import (
	"database/sql"
	"fmt"
)

type SQLTaskRepository struct {
	db *sql.DB
}

func NewSQLTaskRepository(db *sql.DB) TaskRepository {
	return &SQLTaskRepository{db}
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
