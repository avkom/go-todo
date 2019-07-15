package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := os.Getenv("POSTGRES_CONNECTION")
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := NewSQLTaskRepository(db)
	taskService := NewTaskService(taskRepository)
	taskController := NewTaskController(taskService)

	http.Handle("/", taskController)
	http.ListenAndServe(":8080", nil)
}
