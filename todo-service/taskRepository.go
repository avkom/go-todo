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
	query := `SELECT id, title, description	FROM tasks`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskData
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *SQLTaskRepository) GetByID(id string) (task TaskData, err error) {
	query := `SELECT id, title, description	FROM tasks WHERE id = $1`
	row := s.db.QueryRow(query, id)
	err = row.Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return task, err
	}
	return task, nil
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
