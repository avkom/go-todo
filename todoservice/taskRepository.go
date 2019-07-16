package main

import (
	"database/sql"
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
	query := `INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, task.ID, task.Title, task.Description)
	return err
}

func (s *SQLTaskRepository) Update(task TaskData) error {
	query := `UPDATE tasks SET title = $2, description = $3 WHERE id = $1`
	_, err := s.db.Exec(query, task.ID, task.Title, task.Description)
	return err
}

func (s *SQLTaskRepository) Delete(id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
